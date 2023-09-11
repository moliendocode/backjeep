package routes

import (
	"backjeep/controllers"
	"backjeep/repositories"
	"backjeep/utils"

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
	r.POST("/upload", utils.UploadImages)
	r.GET("items", controllers.GetItems)
}
