package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var cld *cloudinary.Cloudinary

func InitCloudinary() {
	InitViper()
	cloudName := viper.GetString("CLOUD_NAME")
	apiKey := viper.GetString("API_KEY")
	apiSecret := viper.GetString("API_SECRET")

	if cloudName == "" || apiKey == "" || apiSecret == "" {
		log.Fatal("Error: Las variables de configuración de Cloudinary no están definidas.")
	}

	var err error
	cld, err = cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Fatalf("Error al inicializar Cloudinary: %s", err)
	}
}

func UploadImages(c echo.Context) ([]string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, fmt.Errorf("error al obtener el formulario: %v", err)
	}

	files := form.File["images"]
	var urls []string

	ctx := context.Background()
	uploadParams := uploader.UploadParams{
		Folder: "campjeep/",
	}

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return nil, fmt.Errorf("error al abrir el archivo: %v", err)
		}

		uploadResult, err := cld.Upload.Upload(ctx, src, uploadParams)
		src.Close()
		if err != nil {
			return nil, fmt.Errorf("error al subir la imagen: %v", err)
		}

		urls = append(urls, uploadResult.SecureURL)
	}

	return urls, nil
}
