package service

import (
	"todo-list/repository"
)

type IUserService interface {
}

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us *UserService) CreateUser() {

}

func (us *UserService) GetAllUser() {

}

func (us *UserService) GetUserById() {

}
func (us *UserService) DeleteUserById() {

}
func (us *UserService) UpdateUserById() {

}
