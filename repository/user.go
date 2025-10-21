package repository

import (
	"fmt"

	"github.com/felipecveiga/task_manager/model"
	"gorm.io/gorm"
)

//go:generate mockgen -source=./user.go -destination=./user_mock.go -package=repository
type Repository interface {
	CreateUserFromDB(vote *model.User) error
	ExistsUserByEmail(email string) (bool, error)
}

type repository struct {
	DB gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &repository{
		DB: *db,
	}
}

func (r *repository) CreateUserFromDB(user *model.User) error {
	err := r.DB.Create(user).Error
	if err != nil {
		return fmt.Errorf("erro ao criar conta no banco de dados: %w", err)
	}

	return nil
}

func (r *repository) ExistsUserByEmail(email string) (bool, error) {
	var user model.User

	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, fmt.Errorf("erro ao buscar usuario por email: %w", result.Error)
	}

	return true, nil
}
