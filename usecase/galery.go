package usecase

import (
	"demonstrasi/models"
	"demonstrasi/models/payload"
	"demonstrasi/repository/database"
	"demonstrasi/util"
	"mime/multipart"
)

func CreateGalery(req *payload.AddGalery, image *multipart.FileHeader) error {
	result, err := util.UploadFile(image)
	if err != nil {
		return err
	}
	galery := models.TbGalery{
		Image: result,
	}
	if err := database.CreateGalery(&galery); err != nil {
		return err
	}
	return nil
}
