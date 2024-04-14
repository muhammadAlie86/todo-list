package repository

import (
	"todo-list/shared/database"
)

type IRoleRepository interface {
}

type RoleRepository struct {
	db database.IDatabase
}

func NewRoleRepository(db database.IDatabase) IRoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (ur *RoleRepository) CreateRole() {

}

func (ur *RoleRepository) GetAllRole() {

}
func (ur *RoleRepository) GetRoleById(roleId string) {

}

func (ur *RoleRepository) DeleteRoleById(roleId string) {

}

func (ur *RoleRepository) UpdateRoleById(roleId string) {

}
