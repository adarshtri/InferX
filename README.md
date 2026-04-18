# InferX 🚀

A high-performance AI Inference System built with **Go** and **C++**.

## 📖 Project Overview
InferX is designed to be a robust, scalable inference server that handles requests asynchronously. It features a Go-based API layer for request handling and a C++ core for optimized batch processing. 

The system uses a **Worker Pool** and **Dynamic Batching** to maximize throughput, bridging Go's concurrency with C++'s raw compute power via a zero-copy Cgo interface.

### Current Status
✅ **Day 18**: Implemented Backpressure and Load Shedding. Verified 1,000+ Requests/sec stability under extreme load by gracefully rejecting excess traffic with HTTP 429.

### Project Structure
- `api/`: Go-based server and load testing suite.
  - `internal/server/`: Core worker pool and batching logic.
  - `internal/engine/`: Cgo wrapper for the C++ library.
  - `pkg/config/`: Configuration management (Env vars & scenarios).
  - `pkg/models/`: Shared request/response types.
- `engine/`: High-performance C++ inference core with metadata-aware processing.
- `lib/`: Containerized static libraries (C++).
- `Dockerfile`: Multi-stage build for the hybrid Go/C++ binary.
- `docker-compose.yml`: Orchestration for scenarios and scaling parameters.

## 🚦 Getting Started

### Prerequisites
- **Docker Desktop**: Required for the hybrid build environment.
- **Go**: 1.23 or higher (for local load testing).
- **Make**: For shortcut commands.

### Running the Server
The simplest way to run the entire integrated system is via Docker:
```bash
make run
```
*This handles C++ compilation, Go linking, and environment setup automatically.*

### Testing the Endpoint
Send a single request to the running container:
```bash
curl -X POST http://localhost:8080/infer \
  -H "Content-Type: application/json" \
  -d '{"model": "llama-3", "prompt": "Hello InferX!"}'
```

### Load Testing
To stress-test the server with concurrent requests from your host machine:
```bash
cd api && go run cmd/loadtest/main.go
```
This utility reports on total throughput (req/sec) and success rates.

## 📅 Roadmap
Detailed progress and architectural walk-throughs can be tracked in:
- [mini-milestones.md](docs/mini-milestones.md)
- [cross_language_integration.md](docs/cross_language_integration.md) (Bridge technical details)
