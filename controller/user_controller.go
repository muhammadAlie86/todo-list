package controller

import (
	"net/http"
	"todo-list/service"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) UpdateUserById(w http.ResponseWriter, r *http.Request) {

}
