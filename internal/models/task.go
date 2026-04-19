package models

import (
	"sync/atomic"
	"time"
)

var lastTaskId atomic.Int64

func getNextID() int64 {
	return lastTaskId.Add(1)
}

type Task struct {
	ID          int64
	UserID      int64
	Title       string
	Description string
	Priority    int

	CreatedAt   time.Time
	DueDate     *time.Time
	IsCompleted bool
}

type TaskOption func(*Task)

func WithPriority(p int) TaskOption {
	return func(t *Task) {
		t.Priority = p
	}
}

func WithDescription(d string) TaskOption {
	return func(t *Task) {
		t.Description = d
	}
}

func WithDueDate(d time.Time) TaskOption {
	return func(t *Task) {
		t.DueDate = &d
	}
}

func NewTask(title string, userID int64, opts ...TaskOption) *Task {
	task := &Task{
		ID:          getNextID(),
		UserID:      userID,
		Title:       title,
		Priority:    1,
		CreatedAt:   time.Now(),
		IsCompleted: false,
	}

	for _, opt := range opts {
		opt(task)
	}

	return task
}
