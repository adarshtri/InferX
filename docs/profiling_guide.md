# InferX: Profiling & Performance Audit Guide

This guide explains how to use Go's built-in `pprof` tool to identify bottlenecks in the InferX engine.

## 🚀 Quick Start (Automated)

While the server is running under load, run:
```bash
make profile
```
This will capture 10 seconds of CPU activity and automatically open a Flame Graph in your browser at `http://localhost:8081`.

## 🛠 Manual Commands

### 1. CPU Profile (Where is the time going?)
Captures a 30-second snapshot of CPU activity:
```bash
go tool pprof -http=:8081 http://localhost:8080/debug/pprof/profile?seconds=30
```

### 2. Heap Profile (Where is the memory going?)
Shows where memory is being allocated:
```bash
go tool pprof -http=:8081 http://localhost:8080/debug/pprof/heap
```

### 3. Goroutine Profile (Why is it stuck?)
Shows what every thread is doing right now (useful for debugging deadlocks or channel contention):
```bash
go tool pprof -http=:8081 http://localhost:8080/debug/pprof/goroutine
```

## 📊 How to read the Flame Graph

- **Width:** Represents total time spent. Wider = more expensive.
- **Top Boxes:** The functions currently executing (the "leaf" functions).
- **Search (Top Right):** Search for "engine" or "batch" to highlight InferX-specific logic.

## 🕵️‍♂️ What to look for in InferX
1. **`runtime.cgocall`**: If this is excessively wide, we might be calling the C++ engine too often (try increasing `BATCH_SIZE`).
2. **`encoding/json`**: If this is wide, consider a faster JSON parser.
3. **`runtime.mallocgc`**: If this is wide, we are creating too many temporary objects (garbage collection pressure).
