# InferX Scaling & Performance Report

This document tracks our observations and lessons learned regarding high-throughput inference scaling.

## 📊 Core Definitions

| Metric | Description | Current Observations |
| :--- | :--- | :--- |
| **Throughput** | Number of requests completed per second (RPS). | We are seeing ~5,000 RPS on raw connectivity, but actual processing throughput is limited by our worker count. |
| **Compute Latency** | Time the worker spends simulating "Inference" work. | Linear model: `100ms + (50ms * BatchSize)`. This simulates GPU processing time. |
| **End-to-End Latency** | Total time from request arrival to response. | Often much higher than compute latency due to **Queueing Delay**. |

## 🧪 Key Observations

### 1. The "Queueing Tax" (Day 13)
We've identified that latency is non-linear under saturation. Even with a static compute cost, end-to-end user latency grows in "waves" corresponding to the number of workers:

- **Wave 1 (Requests 1-20)**: ~350ms latency (No wait, just work).
- **Wave 2 (Requests 21-40)**: ~700ms latency (350ms wait + 350ms work).
- **Wave 3 (Requests 41-60)**: ~1050ms latency (700ms wait + 350ms work).

**Conclusion**: To keep latency low, we must either increase workers or increase batching efficiency to drain the queue faster than it fills.

### 2. Batching Efficiency vs. Latency Tradeoff
- **Size Batching**: Improves throughput (processes more per worker) but can increase latency for the "first" person in the batch if traffic is low.
- **Time Batching**: Guarantees a "maximum wait time" (our 100ms timeout) but might result in smaller, less efficient batches under low load.

---

## 📈 Current Scaling Parameters

- **Workers**: 4
- **Batch Size**: 5
- **Batch Timeout**: 100ms
- **Queue Buffer**: 100

*Next: We will begin experimenting with these parameters on Day 14 to find the "Sweet Spot" for maximum throughput with acceptable latency.*
