# Kubernetes Pod Log Aggregator

Go tool that pulls logs from multiple pods concurrently using errgroup + bounded worker pool.

## What it does
- Lists all pods in a namespace (`test-logs`)
- Fetches last 100 lines of logs from each pod in parallel
- Limits maximum concurrent workers to 10 (semaphore)
- Uses `errgroup.WithContext` so any failure cancels all other requests (fail-fast)
- Returns logs + error summary

## How to Run
```bash
cd k8s-pod-log-aggregator
go run .