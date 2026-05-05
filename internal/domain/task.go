package domain

import "time"

type Task struct {
	Id          uint64
	UserId      uint64
	Description string
	Title       string
	Status      TaskStatus
	DeadLine    *time.Time
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}

type TaskStatus string

const (
	NewTaskSatus         TaskStatus = "NEW"
	InProgressTaskStatus TaskStatus = "IN_PROGRESS"
	DoneTaskStatus       TaskStatus = "DONE"
)
