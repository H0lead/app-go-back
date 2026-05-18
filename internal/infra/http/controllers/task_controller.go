package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type TaskController struct {
	taskService app.TaskService
}

func NewTaskController(ts app.TaskService) TaskController {
	return TaskController{
		taskService: ts,
	}
}

func (c TaskController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)

		task, err := requests.Bind(r, requests.TaskRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskController.Save(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		task.UserId = user.Id
		task.Status = domain.NewTaskSatus

		task, err = c.taskService.Save(task)
		if err != nil {
			log.Printf("TaskController.Save(requests.Bind): %s", err)
			InternalServerError(w, err)
			return
		}

		taskDTO := resources.TaskDTO{}
		taskDTO = taskDTO.DomainToDTO(task)

		Success(w, taskDTO)
	}
}

func (c TaskController) FindList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)

		//Todo: add filters for find list and sorting cryteria
		tasks, err := c.taskService.FindList(user.Id)
		if err != nil {
			log.Printf("TaskController.FindList(c.taskService.FindList): %s", err)
			InternalServerError(w, err)
			return
		}

		taskDTO := resources.TaskDTO{}
		tasksDTO := taskDTO.DomainToDTOCollection(tasks)

		Success(w, tasksDTO)
	}
}

func (c TaskController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		task := r.Context().Value(TaskKey).(domain.Task)

		if task.UserId != user.Id {
			Forbidden(w, errors.New("Access denied"))
			return
		}

		taskDTO := resources.TaskDTO{}
		taskDTO = taskDTO.DomainToDTO(task)

		Success(w, taskDTO)
	}
}

//Todo: add method to change (update) Task status

func (c TaskController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		task := r.Context().Value(TaskKey).(domain.Task)

		if task.UserId != user.Id {
			Forbidden(w, errors.New("Access denied"))
			return
		}

		updTask, err := requests.Bind(r, requests.TaskRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskController.Update(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		task.Title = updTask.Title
		task.Description = updTask.Description
		task.DeadLine = &task.UpdatedDate

		task, err = c.taskService.Update(task)
		if err != nil {
			log.Printf("TaskController.Update(c.taskService.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		taskDTO := resources.TaskDTO{}
		taskDTO = taskDTO.DomainToDTO(task)

		Success(w, taskDTO)
	}
}

func (c TaskController) UpdateStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		task := r.Context().Value(TaskKey).(domain.Task)

		if task.UserId != user.Id {
			Forbidden(w, errors.New("Access denied"))
			return
		}

		updTask, err := requests.Bind(r, requests.TaskStatusRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskController.UpdateStatus(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		task.Status = updTask.Status

		task, err = c.taskService.Update(task)
		if err != nil {
			log.Printf("TaskController.Update(c.taskService.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		taskDTO := resources.TaskDTO{}
		taskDTO = taskDTO.DomainToDTO(task)

		Success(w, taskDTO)
	}
}

func (c TaskController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		task := r.Context().Value(TaskKey).(domain.Task)

		if task.UserId != user.Id {
			Forbidden(w, errors.New("Access denied"))
			return
		}

		err := c.taskService.Delete(task.Id)

		if err != nil {
			log.Printf("TaskController.Delete(c.taskService.Delete): %s", err)
		}

		noContent(w)
	}
}
