package router

import (
	"pretty/config"
	"pretty/repo"
	"pretty/services"

	"pretty/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() {
	router := gin.New()
	// port :=
	db := config.ConnectToDB()

	router.Use(gin.Logger())
	defer db.Close()
	userRepo := repo.CreateUserRepoImpl(db)
	userService := services.CreateUserServiceImpl(userRepo)
	controllers.CreateUserController(userService)

}
