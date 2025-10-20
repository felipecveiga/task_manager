package repository

import (
	"fmt"

	"github.com/felipecveiga/task_manager/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTaskFromDB(userID int, task *model.Task) error
	 GetTasksFromDB(userID int) ([]model.Task, error)
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
