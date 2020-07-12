package services

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"simple/pkg/models"
	"simple/pkg/repositories"
	"time"
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

func (s UserService) GenerateUser(r http.Request) (models.User, error) {
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

func (s UserService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s UserService) generateToken(user models.User) (string, error){
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil{
		return "", err
	}
	return t, nil
}


func (s UserService) AuthUser(requestUser models.User) (string, error) {
	user, err := s.repo.GetUserByEmail(requestUser.Email)
	if err != nil{
		return "", err
	}
	if s.CheckPasswordHash(requestUser.Password, user.Password){
		token, err := s.generateToken(user)
		return token, err
	}
	return "", nil

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

