# 🚀 AI Inference System — Day-wise Milestones Table (with Completion Tracker)

| Day | Phase | Build Tasks | Concepts to Learn | Expected Output | Status |
|-----|------|------------|------------------|----------------|--------|
| 1 | Service Skeleton | Create Go HTTP server, `/infer` endpoint, return dummy response | Go basics, HTTP handlers | API runs locally | ✅ |
| 2 | Request Handling | Accept JSON input, parse into struct | Structs, JSON encoding/decoding | API accepts structured input | ✅ |
| 3 | Queue | Add buffered channel, push requests into queue | Channels (buffered vs unbuffered) | Requests queued asynchronously | ✅ |
| 4 | Worker | Add single goroutine consuming queue | Goroutines, producer-consumer | Async processing works | ✅ |
| 5 | Worker Pool | Add multiple workers (4–8) | Concurrency patterns, worker pools | Parallel processing | ✅ |
| 6 | Load Test | Create simple load test script | Load testing basics | System tested under load | ⬜ |
| 7 | Cleanup | Refactor code, improve structure, add logs | Code organization | Stable base system | ⬜ |
| 8 | Batching | Collect requests into slice | Slices, grouping logic | Basic batching works | ⬜ |
| 9 | Time Batching | Flush batch every X ms | Timers (`time.Ticker`) | Time-based batching | ⬜ |
| 10 | Size Batching | Flush when batch size reached | Conditional triggers | Size-based batching | ⬜ |
| 11 | Combined Batching | Flush on timeout OR size | `select` statement | Dynamic batching | ⬜ |
| 12 | Inference Sim | Add compute delay per batch | Workload simulation | Simulated inference | ⬜ |
| 13 | Latency | Measure request latency | Latency basics | Latency logging | ⬜ |
| 14 | Experiments | Tune batch size, workers | Latency vs throughput | Performance insights | ⬜ |
| 15 | C++ Setup | Create simple C++ batch function | Compilation basics | C++ module ready | ⬜ |
| 16 | Integration | Call C++ from Go (cgo) | FFI basics | Go ↔ C++ working | ⬜ |
| 17 | Optimization | Optimize C++ batch processing | Memory layout, avoiding copies | Faster execution | ⬜ |
| 18 | Backpressure | Add bounded queue, reject overload | Load shedding | Stable under load | ⬜ |
| 19 | Metrics | Track latency, throughput, queue size | Observability basics | Metrics available | ⬜ |
| 20 | Profiling | Identify bottlenecks | Performance analysis | Bottleneck insights | ⬜ |
| 21 | Finalize | Clean code, add README, diagram | System design clarity | Portfolio-ready project | ⬜ |

---

# ✅ Legend
- ⬜ Not started  
- 🟨 In progress  
- ✅ Completed  

---

# ⚠️ Rule

> Do not move to next day until current milestone works end-to-end.