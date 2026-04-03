## HTTP Metrics Server

A lightweight Go-based server that provides system health status and request metrics. This project demonstrates foundational concepts in building web servers, including request routing, JSON encoding, and thread-safe state management.

## Features

- **Health Monitoring**: A dedicated endpoint to check system status and uptime.
- **Metrics Tracking**: Real-time tracking of total requests processed by the server.
- **Concurrent Safety**: Uses `sync/atomic` for thread-safe counter updates, ensuring accuracy under high load.
- **Robust Error Handling**: Implements proper HTTP status codes for method validation and internal encoding errors.

## Tech Stack

- **Language**: Go  
- **Libraries**: Standard library only  
  - `net/http` – Server implementation and routing  
  - `encoding/json` – JSON response formatting  
  - `sync/atomic` – Thread-safe memory operations  
  - `time` – Uptime calculation and tracking

## Concepts Learned

* **Server State Management**: Storing data (like `startTime`) within a struct to maintain context across different request handlers.
* **Atomic Operations**: Implementing `sync/atomic` to prevent "race conditions" when multiple users hit the server at the exact same time.
* **Middleware Logic**: Manually checking `r.Method` to ensure endpoints only respond to authorized HTTP verbs (GET).
* **JSON Streaming**: Using `json.NewEncoder(w).Encode()` to write responses directly to the network socket efficiently.

## Endpoints

| Endpoint | Method | Response Type | Description |
| :--- | :--- | :--- | :--- |
| `/health` | `GET` | `application/json` | Returns system status and uptime in seconds. |
| `/metrics` | `GET` | `text/plain` | Returns a plain-text count of total requests handled. |

## How to Run

```bash
# Save the code as main.go
go run main.go

# To test the endpoints (in a new terminal):
curl http://localhost:8080/health
curl http://localhost:8080/metrics
```