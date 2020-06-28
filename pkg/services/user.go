package services

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"simple/pkg/models"
	"simple/pkg/repositories"
)

type UserService struct {
	repo *repositories.UsersRepository
}

func NewUserService( repo *repositories.UsersRepository ) *UserService {
return &UserService{repo:repo}
}

func (s UserService) generateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s UserService) GenerateUser(w http.ResponseWriter, r http.Request) (models.User, error) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		return user, err
	}
	return user, nil
}

func (s UserService) UserExists(email string) (bool, error){
	cnt, err := s.repo.UserCount(email)
	if err != nil{
		return false, nil
	}
	if cnt > 0{
		return true, nil
	}
	return false, nil
}

func (s *UserService)GetAll() []models.User {
	return s.repo.FetchAll()
}

func (s *UserService)InsertUser(user models.User) error {
	hashPassword, err := s.generateHashPassword(user.Password)
	if err != nil{
		return err
	}
	user.Password = hashPassword
	err = s.repo.Store(user)
	if err != nil{
		return err
	}
	return nil
}

