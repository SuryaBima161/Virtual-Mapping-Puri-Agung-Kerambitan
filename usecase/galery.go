package usecase

import (
	"demonstrasi/models"
	"demonstrasi/models/payload"
	"demonstrasi/repository/database"
	"demonstrasi/util"
	"fmt"
	"mime/multipart"
)

func CreateGalery(req *payload.AddGalery, image *multipart.FileHeader) error {
	result, err := util.UploadFile(image)
	if err != nil {
		return err
	}
	galery := models.TbGalery{
		Id_Information: req.InformationID,
		Image:          result,
	}
	if err := database.CreateGalery(&galery); err != nil {
		return err
	}
	return nil
}
func GetGalery() (resp []payload.GetGaleryRespone, err error) {
	inf, err := database.GetGalery()
	if err != nil {
		return []payload.GetGaleryRespone{}, err
	}
	resp = make([]payload.GetGaleryRespone, len(inf))
	for i, data := range inf {
		resp[i] = payload.GetGaleryRespone{
			Image: data.Image,
		}
	}

	return resp, nil
}

func GetGaleryByRating() ([]payload.GetHomePageRespone, error) {
	galeries, err := database.GetGaleryWithAvgRating()
	if err != nil {
		return nil, fmt.Errorf("failed to get galery with average rating: %v", err)
	}

	resp := make([]payload.GetHomePageRespone, len(galeries))
	for i, data := range galeries {
		resp[i] = payload.GetHomePageRespone{
			ID:     data.ID,
			Image:  data.Image,
			Rating: data.Rating,
		}
	}

	return resp, nil
}
