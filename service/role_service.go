package service

import (
	"todo-list/repository"
)

type IRoleService interface {
}

type RoleService struct {
	roleRepository repository.IRoleRepository
}

func NewRoleService(roleRepository repository.IRoleRepository) IRoleService {
	return &RoleService{
		roleRepository: roleRepository,
	}
}

func (us *RoleService) CreateRole() {

}

func (us *RoleService) GetAllRole() {

}

func (us *RoleService) GetRoleById() {

}

func (us *RoleService) DeleteRoleById() {

}

func (us *RoleService) UpdateRoleById() {

}
