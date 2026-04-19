package models

import "time"

type Task struct {
	ID          int
	Title       string
	Description string
	Priority    int

	CreatedAt   time.Time
	ExpiredAt   time.Time
	DueDate     time.Time
	IsCompleted bool
}
