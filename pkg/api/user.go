package api

import (
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
