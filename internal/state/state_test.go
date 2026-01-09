package state

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// TestStateStruct verifies the State struct has all required fields with correct JSON tags
func TestStateStruct(t *testing.T) {
	state := State{
		StoryKey:    "PROJ-123",
		StepIndex:   2,
		TotalSteps:  5,
		StartStatus: "Todo",
	}

	// Verify JSON serialization includes all fields
	data, err := json.Marshal(state)
	if err != nil {
		t.Fatalf("failed to marshal state: %v", err)
	}

	var decoded State
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal state: %v", err)
	}

	if decoded.StoryKey != "PROJ-123" {
		t.Errorf("StoryKey: got %q, want %q", decoded.StoryKey, "PROJ-123")
	}
	if decoded.StepIndex != 2 {
		t.Errorf("StepIndex: got %d, want %d", decoded.StepIndex, 2)
	}
	if decoded.TotalSteps != 5 {
		t.Errorf("TotalSteps: got %d, want %d", decoded.TotalSteps, 5)
	}
	if decoded.StartStatus != "Todo" {
		t.Errorf("StartStatus: got %q, want %q", decoded.StartStatus, "Todo")
	}
}

// TestSaveWritesValidJSON verifies Save writes valid JSON to file
func TestSaveWritesValidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	mgr := NewManager(tmpDir)

	state := State{
		StoryKey:    "PROJ-456",
		StepIndex:   1,
		TotalSteps:  3,
		StartStatus: "In Progress",
	}

	if err := mgr.Save(state); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	// Read the file directly and verify it's valid JSON
	data, err := os.ReadFile(filepath.Join(tmpDir, ".bmad-state.json"))
	if err != nil {
		t.Fatalf("failed to read state file: %v", err)
	}

	var decoded State
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("file contains invalid JSON: %v", err)
	}

	if decoded != state {
		t.Errorf("saved state mismatch: got %+v, want %+v", decoded, state)
	}
}

// TestLoadReturnsSavedState verifies Load returns previously saved state
func TestLoadReturnsSavedState(t *testing.T) {
	tmpDir := t.TempDir()
	mgr := NewManager(tmpDir)

	original := State{
		StoryKey:    "PROJ-789",
		StepIndex:   0,
		TotalSteps:  4,
		StartStatus: "Todo",
	}

	if err := mgr.Save(original); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	loaded, err := mgr.Load()
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if loaded != original {
		t.Errorf("loaded state mismatch: got %+v, want %+v", loaded, original)
	}
}

// TestLoadReturnsErrNoStateWhenFileMissing verifies Load returns ErrNoState when file doesn't exist
func TestLoadReturnsErrNoStateWhenFileMissing(t *testing.T) {
	tmpDir := t.TempDir()
	mgr := NewManager(tmpDir)

	_, err := mgr.Load()
	if err == nil {
		t.Fatal("expected error when file missing, got nil")
	}

	if !errors.Is(err, ErrNoState) {
		t.Errorf("expected ErrNoState, got %v", err)
	}
}

// TestLoadReturnsErrorForInvalidJSON verifies Load returns error for malformed JSON
func TestLoadReturnsErrorForInvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	mgr := NewManager(tmpDir)

	// Write invalid JSON directly
	statePath := filepath.Join(tmpDir, ".bmad-state.json")
	if err := os.WriteFile(statePath, []byte("{invalid json}"), 0644); err != nil {
		t.Fatalf("failed to write invalid JSON: %v", err)
	}

	_, err := mgr.Load()
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}

	// Should NOT be ErrNoState - it's a different kind of error
	if errors.Is(err, ErrNoState) {
		t.Error("should not return ErrNoState for invalid JSON")
	}
}

// TestClearRemovesExistingFile verifies Clear removes the state file
func TestClearRemovesExistingFile(t *testing.T) {
	tmpDir := t.TempDir()
	mgr := NewManager(tmpDir)

	// Save state first
	state := State{StoryKey: "PROJ-123", StepIndex: 0, TotalSteps: 1, StartStatus: "Todo"}
	if err := mgr.Save(state); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	// Verify file exists
	statePath := filepath.Join(tmpDir, ".bmad-state.json")
	if _, err := os.Stat(statePath); os.IsNotExist(err) {
		t.Fatal("state file should exist after Save")
	}

	// Clear should succeed
	if err := mgr.Clear(); err != nil {
		t.Fatalf("Clear failed: %v", err)
	}

	// File should be gone
	if _, err := os.Stat(statePath); !os.IsNotExist(err) {
		t.Error("state file should not exist after Clear")
	}
}

// TestClearIsIdempotent verifies Clear returns nil when file doesn't exist
func TestClearIsIdempotent(t *testing.T) {
	tmpDir := t.TempDir()
	mgr := NewManager(tmpDir)

	// Clear on non-existent file should succeed
	if err := mgr.Clear(); err != nil {
		t.Errorf("Clear should be idempotent, got error: %v", err)
	}

	// Clear again should still succeed
	if err := mgr.Clear(); err != nil {
		t.Errorf("Clear should be idempotent on second call, got error: %v", err)
	}
}

// TestExistsReturnsTrueWhenFileExists verifies Exists returns true when state file exists
func TestExistsReturnsTrueWhenFileExists(t *testing.T) {
	tmpDir := t.TempDir()
	mgr := NewManager(tmpDir)

	// Save state first
	state := State{StoryKey: "PROJ-123", StepIndex: 0, TotalSteps: 1, StartStatus: "Todo"}
	if err := mgr.Save(state); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	if !mgr.Exists() {
		t.Error("Exists should return true when file exists")
	}
}

// TestExistsReturnsFalseWhenFileMissing verifies Exists returns false when no state file
func TestExistsReturnsFalseWhenFileMissing(t *testing.T) {
	tmpDir := t.TempDir()
	mgr := NewManager(tmpDir)

	if mgr.Exists() {
		t.Error("Exists should return false when file missing")
	}
}
