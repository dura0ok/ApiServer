package main

import (
	"fmt"
	"net/http"
	"simple/pkg/api"
	"simple/pkg/repositories"
	"simple/pkg/services"
)

func main() {
	//==== Building Services ==== ///
	userService := services.NewUserService(repositories.NewUsersRepository())

	//==== Building Controllers ==== ///
	userController := api.UserController{UserService: *userService}

	fmt.Println("Server is started....")
	http.HandleFunc("/api/users", userController.GetUsers)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}