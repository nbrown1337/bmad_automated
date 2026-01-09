// Package state provides lifecycle execution state persistence for resume functionality.
//
// When a lifecycle execution fails (e.g., due to a Claude CLI error), the state
// is saved to disk so that execution can be resumed from the point of failure
// rather than starting over from the beginning. This is particularly valuable
// for long-running story lifecycles.
//
// Key types:
//   - [State] represents the persisted execution state (story key, step index, etc.)
//   - [Manager] handles state persistence operations (save, load, clear)
//
// The state file is stored as a hidden JSON file ([StateFileName]) in the working
// directory. State is written atomically using a temp file and rename pattern
// to prevent corruption on crash.
package state

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// StateFileName is the name of the state file in the working directory.
// It is a hidden file (prefixed with ".") to avoid cluttering the directory.
// The file contains JSON-encoded [State] data.
const StateFileName = ".bmad-state.json"

// ErrNoState is a sentinel error returned by [Manager.Load] when no state file
// exists. Callers should treat this as a signal to start a fresh execution
// rather than as an error condition.
var ErrNoState = errors.New("no state file exists")

// State represents the persisted lifecycle execution state.
//
// This struct is serialized to JSON and saved to disk when a lifecycle
// execution fails, enabling resume from the point of failure.
type State struct {
	// StoryKey is the identifier of the story being processed.
	StoryKey string `json:"story_key"`

	// StepIndex is the 0-based index of the step that failed or is next to execute.
	// On resume, execution continues from this step.
	StepIndex int `json:"step_index"`

	// TotalSteps is the total number of steps in the lifecycle sequence.
	// Used for progress display and validation.
	TotalSteps int `json:"total_steps"`

	// StartStatus is the story's status when execution began.
	// Stored for debugging and context when viewing saved state.
	StartStatus string `json:"start_status"`
}

// Manager handles state persistence operations.
//
// The Manager uses a directory-based approach where the state file is stored
// in a configurable directory. This enables testability by allowing tests to
// use temporary directories.
type Manager struct {
	// dir is the working directory where the state file is stored.
	dir string
}

// NewManager creates a new state manager for the given directory.
//
// The dir parameter specifies the working directory where the state file
// ([StateFileName]) will be stored. Pass "." for the current directory,
// or a temp directory for testing.
func NewManager(dir string) *Manager {
	return &Manager{dir: dir}
}

// statePath returns the full path to the state file.
func (m *Manager) statePath() string {
	return filepath.Join(m.dir, StateFileName)
}

// Save persists the state to disk atomically.
//
// The state is first written to a temporary file, then renamed to the final
// location. This temp file + rename pattern ensures crash safety: the state
// file is either fully written or not present, never corrupted.
func (m *Manager) Save(state State) error {
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	// Write to temp file first for atomic operation
	tmpPath := m.statePath() + ".tmp"
	if err := os.WriteFile(tmpPath, data, 0644); err != nil {
		return err
	}

	// Rename temp file to final location (atomic on POSIX)
	return os.Rename(tmpPath, m.statePath())
}

// Load reads the state from disk.
//
// Returns [ErrNoState] if no state file exists, indicating the caller should
// start a fresh execution. Returns other errors for read or parse failures.
func (m *Manager) Load() (State, error) {
	data, err := os.ReadFile(m.statePath())
	if err != nil {
		if os.IsNotExist(err) {
			return State{}, ErrNoState
		}
		return State{}, err
	}

	var state State
	if err := json.Unmarshal(data, &state); err != nil {
		return State{}, err
	}

	return state, nil
}

// Clear removes the state file if it exists.
//
// This should be called after successful lifecycle completion to clean up.
// The method is idempotent: calling Clear when no state file exists is not
// an error and returns nil.
func (m *Manager) Clear() error {
	err := os.Remove(m.statePath())
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

// Exists returns true if a state file exists.
//
// This is a quick check that can be used to determine if there is saved state
// to resume from, without loading and parsing the full state data.
func (m *Manager) Exists() bool {
	_, err := os.Stat(m.statePath())
	return err == nil
}
