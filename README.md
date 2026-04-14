# InferX 🚀

A high-performance AI Inference System built with **Go** and **C++**.

## 📖 Project Overview
InferX is designed to be a robust, scalable inference server that handles requests asynchronously. It features a Go-based API layer for request handling and a C++ core for optimized batch processing.

This project is being developed following a structured **21-day roadmap**, moving from a simple service skeleton to a fully optimized, multi-worker inference system.

## 🛠 Project Structure
- **/api/cmd/server**: Application entry point.
- **/api/pkg/handlers**: HTTP request handlers and business logic.
- **/api/pkg/models**: Global data structures and types.
- **/docs**: Milestone tracking and design documentation.
- **README.md**: Project overview and instructions.

## 🚀 Current Status: Day 2 Complete
- [x] Initialized Go module.
- [x] Implemented `/infer` endpoint.
- [x] **New**: Structured project into `cmd` and `pkg` layout.
- [x] **New**: Implemented JSON decoding for inference requests.

## 🚦 Getting Started

### Prerequisites
- **Go**: 1.26.2 or higher.

### Running the API
To start the inference service:
```bash
cd api
go run cmd/server/main.go
```

### Testing the Endpoint
You can now send structured JSON payloads to the service:
```bash
curl -X POST http://localhost:8080/infer \
  -H "Content-Type: application/json" \
  -d '{"model": "llama-3", "prompt": "Hello InferX!"}'
```

Expected output:
```json
{"status":"received","message":"Inference request successfully parsed","model":"llama-3"}
```

## 📅 Roadmap
Detailed progress can be tracked in [mini-milestones.md](docs/mini-milestones.md).
