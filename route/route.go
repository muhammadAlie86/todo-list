package route

import (
	"todo-list/controller"

	"github.com/go-chi/chi/v5"
)

func SetupRouterRole(r *chi.Mux, userControler *controller.RoleController) *chi.Mux {

	return r

}
func SetupRouterUser(r *chi.Mux, userControler *controller.UserController) *chi.Mux {

	return r

}
