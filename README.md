<p align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/b/bc/Hermes_Ingenui_%28Vatican%29.jpg/220px-Hermes_Ingenui_%28Vatican%29.jpg" alt="Hermes" width="120"/>
</p>

<h1 align="center">Hermes</h1>

<p align="center">
  <strong>A hybrid C++/Go vector similarity search engine</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white" alt="Go"/>
  <img src="https://img.shields.io/badge/C++-11+-00599C?logo=cplusplus&logoColor=white" alt="C++"/>
  <img src="https://img.shields.io/badge/License-MIT-green" alt="License"/>
</p>

---

## Overview

**Hermes** is a lightweight, high-performance vector database built with a hybrid architecture:

- 🚀 **Go** handles the HTTP server, concurrency, and API layer
- ⚡ **C++** powers the computationally intensive distance calculations via CGO

This approach combines Go's simplicity for web services with C++'s raw performance for numerical operations—a pattern used in production systems like TensorFlow and PyTorch.

---

## Features

| Feature | Description |
|---------|-------------|
| **REST API** | Simple `/insert` and `/search` endpoints |
| **Thread-Safe** | Concurrent-safe storage using `sync.RWMutex` |
| **CGO Bridge** | C++ distance calculations called directly from Go |
| **K-NN Search** | Find the k nearest neighbors to any query vector |
| **Latency Tracking** | Built-in timing for performance monitoring |

---

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                      Python Client                          │
│                       (client.py)                           │
└─────────────────────────┬───────────────────────────────────┘
                          │ HTTP (JSON)
                          ▼
┌─────────────────────────────────────────────────────────────┐
│                       Go Server                             │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │   main.go   │  │  store.go   │  │     bridge.go       │  │
│  │  HTTP API   │──│  In-Memory  │──│   CGO Interface     │  │
│  │  Handlers   │  │   Storage   │  │                     │  │
│  └─────────────┘  └─────────────┘  └──────────┬──────────┘  │
└───────────────────────────────────────────────┼─────────────┘
                                                │ CGO FFI
                                                ▼
┌─────────────────────────────────────────────────────────────┐
│                     C++ Math Engine                         │
│  ┌──────────────────┐  ┌──────────────────────────────────┐ │
│  │  vector_math.h   │  │           math.cpp               │ │
│  │   Header File    │  │  euclidean_distance() impl       │ │
│  └──────────────────┘  └──────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────┘
```

---

## Quick Start

### Prerequisites

- Go 1.21+
- GCC/Clang with C++11 support
- Python 3.8+ (for the test client)

### Run the Server

```bash
# Clone the repository
git clone https://github.com/yourusername/hermes.git
cd hermes

# Build and run (CGO compiles the C++ automatically)
go run .
```

You should see:
```
Hermes
A hybrid C++ and Go
```

The server is now running on `http://localhost:8080`.

### Test with the Python Client

```bash
# Install requests if needed
pip install requests

# Run the test client
python client.py
```

This will:
1. Insert 50,000 random 1536-dimensional vectors
2. Perform a k-NN search and display results with latency

---

## API Reference

### Insert a Vector

```bash
POST /insert
Content-Type: application/json

{
  "id": "doc_123",
  "vector": [0.1, 0.2, 0.3, ...]
}
```

**Response:** `201 Created`

### Search for Similar Vectors

```bash
POST /search
Content-Type: application/json

{
  "vector": [0.1, 0.2, 0.3, ...],
  "k": 5
}
```

**Response:**
```json
{
  "results": [
    {"id": "doc_456", "distance": 0.234},
    {"id": "doc_789", "distance": 0.567}
  ],
  "count": 2,
  "latency": "1.234ms"
}
```

---

## Project Structure

```
hermes/
├── main.go          # HTTP server and route handlers
├── store.go         # Thread-safe in-memory vector storage
├── bridge.go        # CGO bridge to C++ functions
├── math.cpp         # C++ Euclidean distance implementation
├── vector_math.h    # C header for cross-language linkage
├── client.py        # Python test client
└── go.mod           # Go module definition
```

---

## How CGO Works Here

The magic happens in `bridge.go`:

```go
/*
#cgo CXXFLAGS: -std=c++11
#include "vector_math.h"
*/
import "C"

func EuclideanDistance(a, b []float32) float32 {
    ptrA := (*C.float)(unsafe.Pointer(&a[0]))
    ptrB := (*C.float)(unsafe.Pointer(&b[0]))
    return float32(C.euclidean_distance(ptrA, ptrB, C.int(len(a))))
}
```

Go passes raw pointers to C++, which performs the math and returns the result. Zero copying, maximum performance.

---

## 📊 Benchmarks

**Hardware:** Apple Silicon (M-Series)  
**Vector Dimension:** 1536 (OpenAI Standard)

| Dataset Size | Operation | Latency (Server Internal) |
|--------------|-----------|---------------------------|
| 1,000 Vectors | Linear Scan | ~1.3 ms |
| 10,000 Vectors | Linear Scan | ~13.5 ms |
| 50,000 Vectors | Linear Scan | ~69.2 ms |

> **Note:** Latency scales linearly O(n) as expected for exact search implementations. Approximate algorithms like HNSW would achieve O(log n) complexity.

---

## Roadmap

- [ ] **HNSW Indexing** — O(log n) approximate nearest neighbor search
- [ ] **Persistence** — Disk-based storage with memory-mapped files
- [ ] **Cosine Similarity** — Additional distance metrics
- [ ] **Batch Operations** — Insert/search multiple vectors at once
- [ ] **Docker Support** — Containerized deployment

---

## License

MIT License — feel free to use this in your own projects!

---

<p align="center">
  <em>Named after Hermes, the Greek god of messengers — because finding the nearest vector should be fast.</em>
</p>
