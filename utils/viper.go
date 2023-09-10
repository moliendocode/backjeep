package utils

import (
	"log"

	"github.com/spf13/viper"
)

func InitViper() {
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("Error al leer el archivo de configuración: %s", err)
		}
	}
}
