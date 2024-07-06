package util

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
)

func UploadFileToCloudinary(file *multipart.FileHeader) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	uploadResult, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{Folder: "vmap"})
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
