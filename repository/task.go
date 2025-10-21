package repository

import (
	"fmt"

	"github.com/felipecveiga/task_manager/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTaskFromDB(userID int, task *model.Task) error
	GetTasksFromDB(userID int) ([]model.Task, error)
	UpdateTaskInDB(userID int, taskID int, updatedTask *model.Task) error
	DeleteTaskFromDB(userID int, taskID int) error
}

type taskRepository struct {
	DB gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		DB: *db,
	}
}

func (r *taskRepository) CreateTaskFromDB(userID int, task *model.Task) error {
	taskUser := &model.Task{
		Title:       task.Title,
		Description: task.Description,
		UserID:      userID,
		Status:      task.Status,
	}

	err := r.DB.Create(taskUser).Error
	if err != nil {
		return fmt.Errorf("erro ao criar tarefa no banco de dados: %w", err)
	}

	return nil
}

func (r *taskRepository) GetTasksFromDB(userID int) ([]model.Task, error) {
	var tasks []model.Task
	err := r.DB.Where("user_id = ? AND deleted_at IS NULL", userID).Find(&tasks).Error
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar tarefas no banco de dados: %w", err)
	}

	return tasks, nil
}

func (r *taskRepository) UpdateTaskInDB(userID int, taskID int, updatedTask *model.Task) error {
	var task model.Task

	err := r.DB.Where("user_id = ? AND id = ? AND deleted_at IS NULL", userID, taskID).First(&task).Error
	if err != nil {
		return fmt.Errorf("erro ao buscar tarefa no banco de dados: %w", err)
	}

	if updatedTask.Title != "" {
		task.Title = updatedTask.Title
	}

	if updatedTask.Description != "" {
		task.Description = updatedTask.Description
	}

	if updatedTask.Status != "" {
		switch updatedTask.Status {
		case model.Pending, model.InProgress, model.Done:
			task.Status = updatedTask.Status
		default:
			return fmt.Errorf("status inválido")
		}
	}

	err = r.DB.Save(&task).Error
	if err != nil {
		return fmt.Errorf("erro ao atualizar tarefa no banco de dados: %w", err)
	}

	return nil
}

func (r *taskRepository) DeleteTaskFromDB(userID int, taskID int) error {
	err := r.DB.Where("user_id = ? AND id = ?", userID, taskID).Delete(&model.Task{}).Error
	if err != nil {
		return fmt.Errorf("erro ao deletar tarefa no banco de dados: %w", err)
	}

	return nil
}
