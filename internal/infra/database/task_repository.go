package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

const TasksTabelName = "tasks"

type task struct {
	Id          uint64            `db:"id,omitempty"`
	UserId      uint64            `db:"user_id"`
	Description string            `db:"description"`
	Title       string            `db:"title"`
	Status      domain.TaskStatus `db:"status"`
	DeadLine    *time.Time        `db:"deadline"`
	CreatedDate time.Time         `db:"created_date"`
	UpdatedDate time.Time         `db:"updated_date"`
	DeletedDate *time.Time        `db:"deleted_date"`
}
