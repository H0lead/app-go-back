package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type TaskStatusRequest struct {
	Status domain.TaskStatus `json:"status" validate:"required"`
}

func (r TaskStatusRequest) ToDomainModel() (interface{}, error) {
	return domain.Task{
		Status: r.Status,
	}, nil
}
