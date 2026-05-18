package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type TaskDTO struct {
	Id          uint64            `json:"id"`
	UserId      uint64            `json:"userId"`
	Description string            `json:"description,omitempty"`
	Title       string            `json:"title"`
	Status      domain.TaskStatus `json:"status"`
	DeadLine    *time.Time        `json:"deadline,omitempty"`
}

func (d TaskDTO) DomainToDTO(t domain.Task) TaskDTO {
	return TaskDTO{
		Id:          t.Id,
		UserId:      t.UserId,
		Description: t.Description,
		Title:       t.Title,
		Status:      t.Status,
		DeadLine:    t.DeadLine,
	}
}

func (d TaskDTO) DomainToDTOCollection(ts []domain.Task) []TaskDTO {
	tasks := make([]TaskDTO, len(ts))

	for i := range ts {
		tasks[i] = d.DomainToDTO(ts[i])
	}

	return tasks
}
