# InferX 🚀

A high-performance AI Inference System built with **Go** and **C++**.

## 📖 Project Overview
InferX is designed to be a robust, scalable inference server that handles requests asynchronously. It features a Go-based API layer for request handling and a C++ core for optimized batch processing.

## 🚀 Current Status: Day 4 Complete
- ✅ **Day 1**: Basic HTTP server skeleton.
- ✅ **Day 1-2**: Basic HTTP server with JSON request parsing.
- ✅ **Day 3**: Integrated a buffered channel (`chan`) as a request queue.
- ✅ **Day 4**: Implemented a background worker goroutine to process tasks with simulated 500ms latency.

## 🛠 Project Structure
- `api/pkg/handlers/worker.go`: Background worker logic.
- `api/cmd/server/main.go`: Application entry point; initializes queue, server, and worker.
- `api/pkg/models/`: Global data structures and JSON types.
- `/docs`: Progress tracking and design documentation.
- **README.md**: Project overview and instructions.

## 🚦 Getting Started

### Prerequisites
- **Go**: 1.26.2 or higher.

### Running the API
To start the inference service:
```bash
go run api/cmd/server/main.go
```

### Testing the Endpoint
You can now send structured JSON payloads to the service:
```bash
curl -X POST http://localhost:8080/infer \
  -H "Content-Type: application/json" \
  -d '{"model": "llama-3", "prompt": "Hello InferX!"}'
```

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
