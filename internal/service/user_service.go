package service

import (
	"github.com/andreybrumatti/crud-init/internal/entity"
	"github.com/andreybrumatti/crud-init/internal/repository"
)


type UserService struct{
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) GetAllUsers() ([]entity.User, error) {
	var result []entity.User

	users, err := us.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	result = append(result, users...)
	return result, nil
}