package service

import (
	"crud/entity"
	"crud/repository"
	"errors"
)

type UserService struct {
	Repository repository.UserRepository
}

func (service UserService) Get(id string) (*entity.User, error) {
	user := service.Repository.FindById(id)
	if user == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return user, nil
	}
}
