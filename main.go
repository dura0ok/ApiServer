package main

import (
	"fmt"
	"net/http"
	"simple/pkg/api"
	"simple/pkg/api/middlewares"
	"simple/pkg/repositories"
	"simple/pkg/services"
)

func main() {
	//==== Building Services ==== ///
	userService := services.NewUserService(repositories.NewUsersRepository())

	//==== Building Controllers ==== ///
	userController := api.UserController{UserService: *userService}

	//

	fmt.Println("Server is started....")
	http.HandleFunc("/api/users", userController.GetUsers)
	http.HandleFunc("/api/users/register", userController.Register)
	http.HandleFunc("/api/users/login", userController.Login)
	http.Handle("/api/users/user", middlewares.RequireAuthentication()(http.HandlerFunc(userController.User)))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}