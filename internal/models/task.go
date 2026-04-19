package models

import "time"

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
