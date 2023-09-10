package main

import (
	"backjeep/routes"
	"backjeep/utils"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	utils.InitViper()
	dbURI := viper.GetString("DATABASE_URI")
	utils.InitDB(dbURI)
	defer utils.DB.Close()

	if err := utils.CreateTable(); err != nil {
		log.Printf("No se pudo crear la tabla: %v", err)
		return
	}

	e := echo.New()
	routes.InitializeRoutes(e)
	if err := e.Start(":3000"); err != nil {
		log.Printf("Error al iniciar el servidor: %v", err)
	}
}
