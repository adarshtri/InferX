# InferX 🚀

A high-performance AI Inference System built with **Go** and **C++**.

## 📖 Project Overview
InferX is designed to be a robust, scalable inference server that handles requests asynchronously. It features a Go-based API layer for request handling and a C++ core for optimized batch processing.

This project is being developed following a structured **21-day roadmap**, moving from a simple service skeleton to a fully optimized, multi-worker inference system.

## 🛠 Project Structure
- **/api**: Go HTTP server and request handling logic.
- **/docs**: Milestone tracking and design documentation.
- **README.md**: Project overview and instructions.

## 🚀 Current Status: Day 1 Complete
- [x] Initialized Go module.
- [x] Implemented `/infer` endpoint.
- [x] Standard library `net/http` implementation.

## 🚦 Getting Started

### Prerequisites
- **Go**: 1.26.2 or higher.

### Running the API
To start the inference service:
```bash
cd api
go run main/main.go
```

### Testing the Endpoint
You can verify the service is running using `curl`:
```bash
curl -X POST http://localhost:8080/infer
```

Expected output:
```json
{"message":"Inference request received (dummy)","status":"queued"}
```

## 📅 Roadmap
Detailed progress can be tracked in [mini-milestones.md](docs/mini-milestones.md).
