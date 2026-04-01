package main

import (
	"context"
	"log"
	"sync"
	"time"
)

type WorkerPool struct {
	queue      *TaskQueue
	wg         sync.WaitGroup
	ctx        context.Context
	cancel     context.CancelFunc
	numWorkers int
}

func NewWorkerPool(q *TaskQueue, workers int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		queue:      q,
		ctx:        ctx,
		cancel:     cancel,
		numWorkers: workers,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for {
		select {
		case <-wp.ctx.Done():
			log.Printf("Worker %d shutting down", id)
			return

		case task := <-wp.queue.tasks:
			task.StartedAt = time.Now()
			task.Status = "processing"
			log.Printf("Worker %d started VM %s (%s)", id, task.ID, task.Type)

			duration := map[VMSize]time.Duration{
				Small:  2 * time.Second,
				Medium: 4 * time.Second,
				Large:  8 * time.Second,
			}[task.Type]

			time.Sleep(duration)

			task.CompletedAt = time.Now()
			task.Status = "completed"
			log.Printf("Worker %d completed VM %s", id, task.ID)
		}

	}
}

func (wp *WorkerPool) Stop() {
	wp.cancel()
	wp.wg.Wait()
}
