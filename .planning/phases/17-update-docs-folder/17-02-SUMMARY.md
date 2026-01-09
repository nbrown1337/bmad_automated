---
phase: 17-update-docs-folder
plan: 02
subsystem: docs
tags: [documentation, user-guide, cli-reference, lifecycle, dry-run]

# Dependency graph
requires:
  - phase: 17-update-docs-folder
    plan: 01
    provides: PACKAGES.md with lifecycle and state package documentation
provides:
  - USER_GUIDE.md with v1.1 lifecycle features
  - CLI_REFERENCE.md with v1.1 CLI changes
  - Full documentation for --dry-run flag
  - Error recovery documentation
affects: []

# Tech tracking
tech-stack:
  added: []
  patterns:
    - Lifecycle routing table with arrow notation (create-story -> dev-story -> ...)
    - Dry run output examples for all lifecycle commands

key-files:
  created: []
  modified:
    - docs/USER_GUIDE.md
    - docs/CLI_REFERENCE.md

key-decisions:
  - "Added new sections after Status-Based Automation for logical flow"
  - "Included dry-run output examples to show exact format users will see"
  - "Added State File section to CLI_REFERENCE for technical reference"

patterns-established:
  - "Lifecycle routing tables show full sequence from status to done"
  - "Dry run flag documented with examples for all lifecycle commands"

issues-created: []

# Metrics
duration: 8min
completed: 2026-01-09
---

# Phase 17 Plan 02: USER_GUIDE.md and CLI_REFERENCE.md v1.1 Documentation Summary

**Update USER_GUIDE.md and CLI_REFERENCE.md for v1.1 features**

## Performance

- **Duration:** 8 min
- **Started:** 2026-01-09
- **Completed:** 2026-01-09
- **Tasks:** 2
- **Files modified:** 2

## Accomplishments

### USER_GUIDE.md Updates

- Updated Status-Based Automation section to explain full lifecycle execution
- Added "Full Lifecycle Execution" section with example output and how-it-works
- Added "Dry Run Mode" section with examples for run, queue, and epic commands
- Added "Error Recovery" section with .bmad-state.json format and resume flow
- Added "Resume Capability" and "State File Location" sections under Error Handling
- Updated "Batch Processing" section to mention full lifecycle and --dry-run
- Updated "Processing an Epic" section to mention full lifecycle and --dry-run

### CLI_REFERENCE.md Updates

- Updated `run` command: added --dry-run flag, Lifecycle Routing table, updated behavior
- Updated `queue` command: added --dry-run flag, full lifecycle behavior, dry run output example
- Updated `epic` command: added --dry-run flag, full lifecycle behavior
- Added "State File" section documenting .bmad-state.json location, format, fields, and lifecycle

## Task Commits

Each task was committed atomically:

1. **Task 1: Update USER_GUIDE.md with v1.1 features** - `893d716` (docs)
2. **Task 2: Update CLI_REFERENCE.md with v1.1 changes** - `5834adf` (docs)

**Plan metadata:** (this commit)

## Files Created/Modified

- `docs/USER_GUIDE.md` - Added 131 lines: lifecycle, dry-run, error recovery documentation
- `docs/CLI_REFERENCE.md` - Added 130 lines: --dry-run flags, lifecycle routing, state file section

## Verification Checklist

- [x] USER_GUIDE.md has Full Lifecycle Execution section
- [x] USER_GUIDE.md has Dry Run Mode section
- [x] USER_GUIDE.md has Error Recovery section
- [x] CLI_REFERENCE.md documents --dry-run flag for run, queue, epic
- [x] CLI_REFERENCE.md explains full lifecycle execution
- [x] CLI_REFERENCE.md documents .bmad-state.json
- [x] No broken markdown links

## Decisions Made

- Placed new USER_GUIDE sections after Status-Based Automation for natural reading flow
- Included concrete dry run output examples to show exact format users will see
- Added State File as a new top-level section in CLI_REFERENCE (parallel to Sprint Status File)
- Used arrow notation (create-story -> dev-story -> ...) for lifecycle sequences consistently

## Deviations from Plan

None - plan executed exactly as written.

## Issues Encountered

None

## Next Phase Readiness

- USER_GUIDE.md and CLI_REFERENCE.md fully document v1.1 features
- Documentation milestone v1.2 progressing
- Ready for any remaining 17-xx plans or final milestone completion

---

_Phase: 17-update-docs-folder_
_Completed: 2026-01-09_
