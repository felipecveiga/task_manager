package service

import (
	"github.com/felipecveiga/task_manager/model"
	"github.com/felipecveiga/task_manager/repository"
)

type TaskService interface {
	CreateTask(userID int, task *model.Task) error
}

type taskService struct {
	Repository repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) TaskService {
	return &taskService{
		Repository: r,
	}
}

func (s *taskService) CreateTask(userID int, task *model.Task) error {

	err := s.Repository.CreateTaskFromDB(userID, task)
	if err != nil {
		return err
	}

	return nil
}
