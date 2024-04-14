package repository

import (
	"todo-list/shared/database"
)

type IUserRepository interface {
}

type UserRepository struct {
	db database.IDatabase
}

func NewUserRepository(db database.IDatabase) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser() {

}

func (ur *UserRepository) GetAllUser() {

}
func (ur *UserRepository) GetUserById(userId string) {

}

func (ur *UserRepository) DeleteUserById(userId string) {

}

func (ur *UserRepository) UpdateUserById(userId string) {

}
