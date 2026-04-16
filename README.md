# InferX 🚀

A high-performance AI Inference System built with **Go** and **C++**.

## 📖 Project Overview
InferX is designed to be a robust, scalable inference server that handles requests asynchronously. It features a Go-based API layer for request handling and a C++ core for optimized batch processing.

## 🚀 Current Status: Day 3 Complete
- ✅ **Day 1**: Basic HTTP server skeleton.
- ✅ **Day 2**: JSON request parsing and validation.
- ✅ **Day 3**: Integrated a buffered channel (`chan`) as a request queue with `202 Accepted` response logic.

## 🛠 Project Structure
- `api/cmd/server/main.go`: Application entry point; initializes queue and server.
- `api/pkg/handlers/`: Contains the `Server` struct and HTTP handler methods.
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

**Note:** Since we haven't added workers yet (Day 4), requests will simply stay in the queue. You can see the queue depth increase in the server logs.

## 📅 Roadmap
Detailed progress can be tracked in [mini-milestones.md](docs/mini-milestones.md).
