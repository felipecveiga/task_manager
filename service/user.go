package service

import (
	"fmt"

	"github.com/felipecveiga/task_manager/model"
	"github.com/felipecveiga/task_manager/repository"
)

type Service interface {
	CreateUser(user *model.User) error
}

type service struct {
	Repository repository.Repository
}

func NewUserService(r repository.Repository) Service {
	return &service{
		Repository: r,
	}
}

func (s *service) CreateUser(user *model.User) error {

	hasEmail, err := s.Repository.ExistsUserByEmail(user.Email)
	if err != nil {
		return err
	}

	if hasEmail {
		return fmt.Errorf("email ja cadastrado")
	}

	err = s.Repository.CreateUserFromDB(user)
	if err != nil {
		return err
	}

	return nil
}
