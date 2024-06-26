package usecase

import (
	"demonstrasi/models"
	"demonstrasi/models/payload"
	"demonstrasi/repository/database"
	"demonstrasi/util"
	"mime/multipart"
)

func CreateMonument(req *payload.MonumentRequest, image *multipart.FileHeader) error {
	result, err := util.UploadFile(image)
	if err != nil {
		return err
	}
	galery := models.TbMonument{
		Id_Information: req.InformationID,
		Image:          result,
	}
	if err := database.CreateMonument(&galery); err != nil {
		return err
	}
	return nil
}
func GetMonument() (resp []payload.GetMonumentRespone, err error) {
	inf, err := database.GetMonument()
	if err != nil {
		return []payload.GetMonumentRespone{}, err
	}
	resp = make([]payload.GetMonumentRespone, len(inf))
	for i, data := range inf {
		resp[i] = payload.GetMonumentRespone{
			Image: data.Image,
		}
	}

	return resp, nil
}
