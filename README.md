# InferX 🚀

A high-performance AI Inference System built with **Go** and **C++**.

## 📖 Project Overview
InferX is designed to be a robust, scalable inference server that handles requests asynchronously. It features a Go-based API layer for request handling and a C++ core for optimized batch processing.

## 🚀 Current Status: Day 7 Complete
- ✅ **Day 1-6**: API implementation, Worker Pool, and Load Testing.
- ✅ **Day 7**: System Hardening (Structured Logging, Config Management, Graceful Shutdown, and Package Restructuring).

## 🛠 Project Structure
- `api/cmd/server/main.go`: Entry point for the main inference server.
- `api/internal/server/`: Core inference engine (Worker Pool, Queue, and Handlers).
- `api/pkg/config/`: Configuration management (Env vars & .env support).
- `api/pkg/models/`: Shared models and JSON types.
- `api/cmd/loadtest/main.go`: Concurrent load-testing utility.
- `/docs`: Progress tracking and design documentation.

## 🚦 Getting Started

### Prerequisites
- **Go**: 1.26.2 or higher.

### Running the API
To start the inference service:
```bash
go run api/cmd/server/main.go
```

### Testing the Endpoint
Send a single request:
```bash
curl -X POST http://localhost:8080/infer \
  -H "Content-Type: application/json" \
  -d '{"model": "llama-3", "prompt": "Hello InferX!"}'
```

### Load Testing
To stress-test the server with concurrent requests:
```bash
go run api/cmd/loadtest/main.go
```
This script sends 100 requests (10 at a time) and reports on the success rate and total throughput (req/sec).

Expected output (Day 3+):
```json
{
  "status": "queued",
  "message": "Inference request successfully queued for processing",
  "model": "llama-3"
}
```

**Note:** As of Day 4, a background worker is now processing these tasks sequentially. Each task takes 500ms. You can monitor progress in the server console logs which show "Starting processing..." and "Completed processing...".

## 📅 Roadmap
Detailed progress can be tracked in [mini-milestones.md](docs/mini-milestones.md).
