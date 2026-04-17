---
trigger: always_on
description: Ensures docs/performance_report.md is updated with scaling insights.
---

# Performance Report Tracking

## Goal
Ensure that every performance insight, latency observation, and throughput finding is recorded in `docs/performance_report.md` to maintain a scientific history of the project's scaling journey.

## Workflow
1. **Analyze Results**: After any major code change (Batching, Workers, Inference Sim) or load test, evaluate the impact on RPS and p99 latency.
2. **Identify Observations**: Distill the key performance finding (e.g., "Increasing workers by 2x reduced queueing latency by 40%").
3. **Update Report**:
    - Add NEW learnings to the "Key Observations" section.
    - Update the "Current Scaling Parameters" section with the latest "Best Known Configuration."
    - Refer to the baseline metrics to show progress or regression.
4. **Synchronized Commits**: Always commit updates to `docs/performance_report.md` at the same time as `docs/mini-milestones.md` and `README.md` at the end of each "Day."

## Verification
- Insights must be backed by data (logs or load test results) to be included in the report.
