package main

import (
	"fmt"
	"log"
	"time"
)

type Scheduler struct {
	queue *TaskQueue
	pool  *WorkerPool
}

func NewScheduler(workers int) *Scheduler {
	q := NewTaskQueue()
	p := NewWorkerPool(q, workers)
	return &Scheduler{queue: q, pool: p}
}

func (s *Scheduler) Provision(size VMSize) {
	task := &Task{
		ID:        fmt.Sprintf("vm-%d", time.Now().UnixNano()%10000),
		Type:      size,
		Status:    "queued",
		CreatedAt: time.Now(),
	}

	s.queue.Enqueue(task)
	log.Printf("New Provision request accepted: %s (%s)", task.ID, size)
}
