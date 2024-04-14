package main

import (
	"log"
	"net/http"
	"todo-list/shared/config"
	"todo-list/shared/database"

	"todo-list/controller"
	"todo-list/repository"
	"todo-list/route"
	"todo-list/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	setupMain()
}

func setupMain() {
	cfg := config.InitConfig()
	db := database.NewDataBase(cfg)

	// Init reposity
	userRepository := repository.NewUserRepository(db)
	roleRepository := repository.NewRoleRepository(db)

	// Init service
	userService := service.NewUserService(userRepository)
	roleService := service.NewRoleService(roleRepository)

	// Init controller
	userController := controller.NewUserController(userService)
	roleController := controller.NewRoleController(roleService)

	// Init router utama
	router := chi.NewRouter()

	// Init subrouter
	userRouter := chi.NewRouter()
	roleRouter := chi.NewRouter()

	// Set up rute untuk setiap subrouter
	route.SetupRouterRole(roleRouter, roleController)
	route.SetupRouterUser(userRouter, userController)

	// Menggunakan subrouter sebagai handler untuk rute tertentu di router utama
	router.Mount("/user", userRouter)
	router.Mount("/management", roleRouter)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server started on port 8080")
	log.Fatal(server.ListenAndServe())
	defer db.CloseDatabase()

}
