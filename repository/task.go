package repository

import (
	"fmt"

	"github.com/felipecveiga/task_manager/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTaskFromDB(userID int, task *model.Task) error
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
