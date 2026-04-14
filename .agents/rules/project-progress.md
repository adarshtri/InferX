---
trigger: always_on
description: Ensures development follows the roadmap in docs/mini-milestones.md
---

# Project Progress Tracking

## Goal
Ensure development follows the structured roadmap in `docs/mini-milestones.md` to prevent skipping steps, feature creep, or deviation from the core objectives.

## Workflow
1. **Initial Check**: At the start of every task or conversation, read `docs/mini-milestones.md` to identify the current active milestone (marked with 🟨) or the next pending milestone (marked with ⬜).
2. **Strict Sequencing**: NEVER suggest or implement features from a later "Day" until all tasks from the current "Day" are completed and marked as ✅.
3. **No Skipping**: If a prerequisite step is unfulfilled, point it out to the user and prioritize it.
4. **Deviation Prevention**: If a user request deviates significantly from the roadmap, flag it as a "Deviation" and ask for confirmation before proceeding.
5. **Update Progress**:
    - When starting a task, update the milestone status to 🟨 in `docs/mini-milestones.md`.
    - When a milestone is verified end-to-end (as per the "Expected Output"), update it to ✅.
    - Always commit the status update to `docs/mini-milestones.md` along with the code changes.

## Verification
- A milestone is only "Complete" (✅) when it works end-to-end as described in the "Expected Output" column of the milestones table.
