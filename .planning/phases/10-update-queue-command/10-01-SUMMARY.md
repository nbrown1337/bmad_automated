---
phase: 10-update-queue-command
plan: 01
subsystem: cli
tags: [lifecycle, queue, tdd, cobra]

# Dependency graph
requires:
  - phase: 07-story-lifecycle-executor
    provides: lifecycle.Executor for running full story lifecycle
  - phase: 09-update-epic-command
    provides: Pattern for lifecycle executor usage in CLI commands
provides:
  - Queue command with full lifecycle execution
  - Consistent behavior across run, queue, and epic commands
  - Removed unused QueueRunner dependency from App struct
affects: [11-error-recovery-resume, 12-dry-run-mode]

# Tech tracking
tech-stack:
  added: []
  patterns:
    - lifecycle.Executor for full story lifecycle
    - Interface-based DI for testability

key-files:
  created: []
  modified:
    - internal/cli/queue.go
    - internal/cli/queue_test.go
    - internal/cli/root.go
    - internal/cli/cli_test.go
    - internal/cli/run_test.go

key-decisions:
  - "Follow epic.go pattern exactly for consistency"
  - "Remove app.Queue field since lifecycle executor is created inline"

patterns-established:
  - "All story-processing commands use lifecycle.Executor"

issues-created: []

# Metrics
duration: 4min
completed: 2026-01-09
---

# Phase 10 Plan 01: Queue Command with Lifecycle Execution Summary

**Queue command now executes full story lifecycle (create->dev->review->commit) for each story before moving to next, consistent with run and epic commands**

## Performance

- **Duration:** 4 min
- **Started:** 2026-01-09T02:33:20Z
- **Completed:** 2026-01-09T02:37:09Z
- **TDD Phases:** RED, GREEN, REFACTOR
- **Files modified:** 5

## TDD Cycle

### RED Phase

Wrote comprehensive tests for queue lifecycle execution:

- 2 backlog stories runs full lifecycle for each (8 workflows total)
- Mixed statuses runs appropriate remaining workflows
- Done story is skipped and continues with others
- Workflow failure mid-lifecycle stops processing (fail-fast)
- All done stories returns success with no workflows
- Single story runs full lifecycle
- Story not found returns error
- Missing sprint-status.yaml returns error

Tests failed because queue command was still using QueueRunner instead of lifecycle executor.

### GREEN Phase

Updated queue.go to use lifecycle.Executor following epic.go pattern exactly:

1. Create lifecycle executor with app dependencies
2. Loop through story keys from args (not from epic discovery)
3. Call executor.Execute() for each story
4. Handle router.ErrStoryComplete for done stories (skip with message)
5. Fail-fast on workflow failure
6. Print success message after each story completes
7. Print total stories processed at end
8. Updated help text to describe full lifecycle behavior

### REFACTOR Phase

Removed unused Queue field from App struct:

- Queue command now creates lifecycle executor inline (like epic)
- Removed app.Queue field from App struct definition
- Removed QueueRunner initialization from NewApp()
- Updated all test setup functions to remove Queue field
- Consistent architecture across run, queue, and epic commands

## Task Commits

1. **RED: Failing tests** - `6e94b05` (test)
2. **GREEN: Implementation** - `cff118c` (feat)
3. **REFACTOR: Cleanup** - `d72a84c` (refactor)

## Files Created/Modified

- `internal/cli/queue.go` - Updated to use lifecycle.Executor instead of QueueRunner
- `internal/cli/queue_test.go` - Complete rewrite with lifecycle execution tests
- `internal/cli/root.go` - Removed Queue field from App struct
- `internal/cli/cli_test.go` - Updated test setup functions
- `internal/cli/run_test.go` - Updated test setup functions

## Decisions Made

- **Follow epic.go pattern exactly:** Consistency across all story-processing commands
- **Remove app.Queue field:** Lifecycle executor is created inline where needed, no need for field

## Deviations from Plan

None - plan executed exactly as written.

## Issues Encountered

None

## Next Phase Readiness

- All story-processing commands (run, queue, epic) now use lifecycle executor
- Consistent behavior: full lifecycle per story, done skip, fail-fast
- Ready for Phase 11: Error Recovery & Resume

---

_Phase: 10-update-queue-command_
_Completed: 2026-01-09_
