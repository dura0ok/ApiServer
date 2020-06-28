package services

import (
	"simple/pkg/models"
	"simple/pkg/repositories"
)

type UserService struct {
	repo *repositories.UsersRepository
}

func NewUserService( repo *repositories.UsersRepository ) *UserService {
return &UserService{repo:repo}
}

func (s *UserService)GetAll() []models.User {
	return s.repo.FetchAll()
}