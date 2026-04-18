# InferX 🚀

A high-performance AI Inference System built with **Go** and **C++**. 

## 📖 Project Overview
InferX is a robust, scalable inference server designed to maximize throughput through **Dynamic Batching** and **Zero-Copy Memory Bridging**. It seamlessly integrates Go's lightweight concurrency with C++'s optimized compute performance.

### Project Status: ✅ COMPLETED
InferX has evolved from a simple sequential Go server into a professional-grade hybrid system capable of sustaining **1,000+ Requests/sec** with adaptive backpressure and real-time observability.

## 🧠 Technical Deep Dive

### 1. Hybrid Concurrency (Go ↔ C++)
- **Worker Pool:** A managed pool of goroutines handles request orchestration.
- **Dynamic Batching:** Automatically bundles requests based on size or timeout to maximize C++ throughput.
- **Cgo Bridge:** Uses `unsafe.SliceData` for zero-copy memory visibility between Go and C++.

### 2. Resilience & Stability
- **Load Shedding:** Adaptive backpressure returns **HTTP 429** during overload to protect server p99 latency.
- **Graceful Shutdown:** Ensures all pending batches are processed before the server exits.

### 3. Observability
- **Prometheus Metrics:** Real-time counters for Accepted, Rejected, and Processed requests.
- **Integrated Profiling:** Automated `pprof` workflow for identifying CPU and Memory hotspots.

## 🏗 Project Structure
- `api/`: Go-based server and load testing suite.
- `engine/`: High-performance C++ inference core.
- `docs/`: Technical guides, architectural summaries, and milestones.
- `Makefile`: Automated build and profiling workflow.

## 🚦 Getting Started

### Prerequisites
- **Docker Desktop**: Required for the hybrid build environment.
- **Go 1.23+**: For local load testing.

### Running the Server
```bash
# Standard Run
make run

# Stress-Test Scenario (Overload)
SCENARIO=overload make run
```

### Monitoring & Profiling
- **Metrics:** `curl http://localhost:8080/metrics`
- **Profilier:** `make profile` (Opens Flame Graph in browser)

## 📅 Roadmap & Documentation
- [Architecture Summary](docs/architecture_summary.md)
- [Cross-Language Bridge Deep Dive](docs/cross_language_integration.md)
- [Profiling & Audit Guide](docs/profiling_guide.md)
- [Mini-Milestones (Day-by-Day)](docs/mini-milestones.md)
