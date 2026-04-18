# InferX: Cross-Language Integration Guide

This document explains the internal mechanics of the Go-C++ bridge used in InferX, focusing on memory safety, performance, and data layout.

## 🌉 The Cgo Bridge
InferX uses **Cgo** to bridge the high-level Go API with the high-performance C++ inference core. 

### 1. Memory Layout: The "Zero-Copy" Strategy
To achieve maximum performance, we avoid copying data when passing it to C++. Instead, we pass pointers to memory managed by Go.

#### Slice Headers & `unsafe.SliceData`
In Go, a slice is a header containing a pointer to the underlying array, a length, and a capacity.
```go
type SliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
}
```
We use `unsafe.SliceData(slice)` to extract the `Data` pointer. This is a zero-cost operation that allows C++ to read the exact same RAM as Go.

### 2. Memory Safety & Lifecycle
Sharing memory across the language boundary introduces risks that the Go Garbage Collector (GC) cannot see.

> [!CAUTION]
> **Pointer Longevity:** Go only guarantees that memory passed to a C function stays valid for the duration of that function call.
> **Rule:** Never store Go pointers in C++ global variables or long-lived heap structures.

### 3. Data Alignment
C++ expects contiguous memory. Go's slices of primitives (int, float64, etc.) are guaranteed to be contiguous, making them perfectly compatible with standard C-style arrays (`int*`).

## 🛠 Build System Integration
Because macOS, Linux, and Windows handle dynamic libraries differently (SDK mismatches), InferX uses a **Multi-Stage Docker Build**:

1.  **Builder Stage (Linux):** Compiles C++ into a static archive (`.a`).
2.  **Linking Stage:** The Go compiler links the `.a` into the final binary.
3.  **Runtime Stage:** A minimal image containing only the self-contained binary.

This bypasses the host OS's linker/SDK issues entirely.

## 🚀 Performance Considerations
- **Cgo Overhead:** Every call from Go to C has a small cost (~50ns). 
- **Batching:** To minimize this overhead, we always batch processing into a single call: `ProcessBatch(data, size)`. It is far more efficient to make one call for 100 items than 100 calls for 1 item.
