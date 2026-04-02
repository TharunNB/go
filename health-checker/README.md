# Concurrent URL Health Checker

A lightweight Go CLI tool to check the health of multiple URLs concurrently using goroutines and channels.

## Features

- Concurrent HTTP requests (fan-out pattern)
- Timeout handling to prevent hanging requests
- Robust error handling (DNS issues, invalid URLs, network failures)
- Clean output with status codes
- Efficient synchronization using sync.WaitGroup

## Tech Stack

- Language: Go  
- Libraries: Standard library only  
  - net/http – HTTP requests  
  - context – request timeouts  
  - sync – concurrency control  

## How to Run

```bash
go run .
