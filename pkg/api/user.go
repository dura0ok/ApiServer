package api

import (
	"errors"
	"log"
	"net/http"
	"simple/pkg/services"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (u UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := u.UserService.GetAll()
	err := services.GoodResponse(w, users)
	if err != nil{
		log.Fatalln(err)
	}
}


func (u UserController) Register(w http.ResponseWriter, r *http.Request) {
	check, err := services.CheckRequestMethod(*r, w, "POST")
	if !check{
		return
	}

	user, err := u.UserService.GenerateUser(*r)
	if err != nil{
		_ = services.BadResponse(w, errors.New("generate user error :(").Error())
		log.Fatalln(err)
		return
	}
	userExists, err := u.UserService.UserExists(user.Email)
	if err != nil{
		log.Fatalln(err)
		return
	}
	if userExists{
		services.BadResponse(w, errors.New("user with this email already exists").Error())
		return
	}
	_ = u.UserService.InsertUser(user)
	err = services.GoodResponse(w, "ok")
	if err != nil{
		log.Fatalln(err)
	}
}

func (u UserController) Login(w http.ResponseWriter, r *http.Request) {
	check, err := services.CheckRequestMethod(*r, w, "POST")
	if !check{
		return
	}
	user, err := u.UserService.GenerateUser(*r)
	if err != nil{
		err := services.BadResponse(w, errors.New("we can't read your data :(").Error())
		if err != nil{
			log.Fatalln(err)
		}
		return
	}
	token, err := u.UserService.AuthUser(user)
	if err != nil{
		log.Fatalln(err)
	}
	if token != ""{
		_ = services.GoodResponse(w, map[string]string{"token": token})
		return
	}
	_ = services.BadResponse(w, "Ошибка авторизации")
}

func (u UserController) User(w http.ResponseWriter, r *http.Request) {
	services.GoodResponse(w, map[string]string{"message": "ok"})

}