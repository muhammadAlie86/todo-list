package controller

import (
	"net/http"
	"todo-list/service"
)

type RoleController struct {
	roleService service.IRoleService
}

func NewRoleController(roleService service.RoleService) *RoleController {
	return &RoleController{
		roleService: roleService,
	}
}

func (uc *UserController) CreateRole(w http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) GetAllRole(w http.ResponseWriter, r *http.Request) {

}
func (uc *UserController) GetRoleById(w http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) DeleteRoleById(w http.ResponseWriter, r *http.Request) {

}
func (uc *UserController) UpdateRoleById(w http.ResponseWriter, r *http.Request) {

}
