package main

type TaskQueue struct {
	tasks chan *Task
}

func NewTaskQueue() *TaskQueue {
	return &TaskQueue{tasks: make(chan *Task, 100)}
}

func (q *TaskQueue) Enqueue(t *Task) { q.tasks <- t }
func (q *TaskQueue) Dequeue() *Task  { return <-q.tasks }
