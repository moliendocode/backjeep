package main

import (
	"backjeep/routes"
	"backjeep/utils"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	utils.InitViper()
	dbURI := viper.GetString("DATABASE_URI")
	utils.InitCloudinary()
	utils.InitDB(dbURI)
	defer utils.DB.Close()

	if err := utils.CreateTable(); err != nil {
		log.Printf("No se pudo crear la tabla: %v", err)
		return
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://localhost:3000", "http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	routes.InitializeRoutes(e)
	if err := e.Start(":3001"); err != nil {
		log.Printf("Error al iniciar el servidor: %v", err)
	}
}
