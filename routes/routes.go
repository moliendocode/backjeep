package routes

import (
	"backjeep/controllers"
	"backjeep/repositories"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func InitializeRoutes(e *echo.Echo) {
	userRepo := &repositories.UserRepo{}
	userController := controllers.NewUserController(userRepo)

	jwtConfig := echojwt.Config{
		SigningKey:    []byte(viper.GetString("JWT_SECRET")),
		TokenLookup:   "cookie:jwt",
		SigningMethod: "HS256",
	}

	e.POST("/login", userController.Login)

	r := e.Group("")
	r.Use(echojwt.WithConfig(jwtConfig))
	r.GET("/users", userController.GetAllUsers)
	r.POST("/user", userController.CreateUser)
	r.GET("/items", controllers.GetItems)
	r.POST("/items", controllers.CreateItem)
	r.GET("/item", controllers.GetItemDetails)
	r.PATCH("/item", controllers.UpdateItem)
}
