# InferX: Architectural Concepts & Learning

This document records the engineering decisions and high-concurrency patterns used to make InferX scale to 1,000+ requests per second.

## ⚡ High-Concurrency Patterns

### 1. Atomic Operations vs. Mutexes
When tracking metrics on the "Hot Path" (code executed for every request), we prioritize `sync/atomic` over `sync/Mutex`.

- **The Problem with Mutexes:** Mutexes are heavyweight "locks." If multiple threads (goroutines) try to increment a counter at the same time, the OS must put some threads to sleep and wake them up later (Context Switching). This overhead becomes the bottleneck at high RPS.
- **The Atomic Solution:** `sync/atomic` uses hardware-level CPU instructions (Compare-and-Swap) to update values without ever locking. It is "Lock-Free," meaning no goroutine ever goes to sleep just to update a counter.
- **Usage in InferX:** We use Atomics for all real-time counters (Successes, Rejections, etc.) to ensure monitoring doesn't slow down inference.

## 📊 Observability Standards

### 1. Prometheus Text Format
While JSON is easy for humans to read, InferX is designed to be compatible with industry-standard monitoring tools like Prometheus.

#### Prometheus Format Anatomy:
```text
# HELP requests_total Total number of requests processed
# TYPE requests_total counter
requests_total{model="gpt-4", status="202"} 1055
```
- **HELP/TYPE:** Metadata for the monitoring system.
- **Counter:** A value that only increases (e.g., total requests).
- **Gauge:** A value that goes up and down (e.g., current queue depth).
- **Labels:** Dimensions (inside `{}`) that allow for powerful filtering (e.g., "Show me rejections only for llama-3").

## 🛡️ Stability Patterns

### 1. Load Shedding (Backpressure)
Instead of letting the server "hang" or crash under extreme load, InferX implements **Load Shedding**.
- **Go Pattern:** Using a `select` statement with a `default` case during channel pushes.
- **Outcome:** If the internal queue is full, the server instantly returns **HTTP 429 Too Many Requests**. This protects the latency of already-accepted requests and ensures the server remains responsive.
