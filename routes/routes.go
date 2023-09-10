package routes

import (
	"backjeep/controllers"
	"backjeep/repositories"

	"github.com/labstack/echo/v4"
)

func InitializeRoutes(e *echo.Echo) {
	userRepo := &repositories.UserRepo{}
	userController := controllers.NewUserController(userRepo)

	e.GET("/users", userController.GetAllUsers)
	e.POST("/user", userController.CreateUser)
	e.POST("/login", userController.Login)
}
