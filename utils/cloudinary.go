package utils

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"sync"

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

	var wg sync.WaitGroup

	type result struct {
		url string
		err error
	}

	resultChan := make(chan result, len(files))

	ctx := context.Background()
	uploadParams := uploader.UploadParams{
		Folder: "campjeep/",
	}

	const maxConcurrentUploads = 10
	semaphore := make(chan struct{}, maxConcurrentUploads)

	for _, file := range files {
		wg.Add(1)
		go func(file *multipart.FileHeader) {
			semaphore <- struct{}{}
			defer func() {
				<-semaphore
			}()
			defer wg.Done()

			src, err := file.Open()
			if err != nil {
				resultChan <- result{err: fmt.Errorf("error al abrir el archivo: %v", err)}
				return
			}

			uploadResult, err := cld.Upload.Upload(ctx, src, uploadParams)

			src.Close()
			if err != nil {
				resultChan <- result{err: fmt.Errorf("error al subir la imagen: %v", err)}
				return
			}

			resultChan <- result{url: uploadResult.SecureURL}
		}(file)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var urls []string
	for r := range resultChan {
		if r.err != nil {
			return nil, r.err
		}
		urls = append(urls, r.url)
	}

	return urls, nil
}
