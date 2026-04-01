package main

import "time"

type VMSize string

const (
	Small  VMSize = "small"
	Medium VMSize = "medium"
	Large  VMSize = "large"
)

type Task struct {
	ID          string
	Type        VMSize
	Status      string
	CreatedAt   time.Time
	StartedAt   time.Time
	CompletedAt time.Time
}
