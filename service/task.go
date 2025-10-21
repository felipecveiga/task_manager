package service

import (
	"github.com/felipecveiga/task_manager/model"
	"github.com/felipecveiga/task_manager/repository"
)

//go:generate mockgen -source=./task.go -destination=./task_mock.go -package=service
type TaskService interface {
	CreateTask(userID int, task *model.Task) error
	GetTasksByID(userID int) ([]model.Task, error)
	UpdateTask(userID int, taskID int, updatedTask *model.Task) error
	DeleteTask(userID int, taskID int) error
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

func (s *taskService) GetTasksByID(userID int) ([]model.Task, error) {

	tasks, err := s.Repository.GetTasksFromDB(userID)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *taskService) UpdateTask(userID int, taskID int, updatedTask *model.Task) error {
	
	err := s.Repository.UpdateTaskInDB(userID, taskID, updatedTask)
	if err != nil {
		return err
	}

	return nil
}	

func (s *taskService) DeleteTask(userID int, taskID int) error {
	err := s.Repository.DeleteTaskFromDB(userID, taskID)
	if err != nil {
		return err
	}

	return nil
}

