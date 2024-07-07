package usecase

import (
	"demonstrasi/models"
	"demonstrasi/models/payload"
	"demonstrasi/repository/database"
	"demonstrasi/util"
	"fmt"
	"mime/multipart"

	uuid "github.com/satori/go.uuid"
)

func CreateGalery(req *payload.AddGalery, image *multipart.FileHeader) error {
	result, err := util.UploadFileToCloudinary(image)
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
			Image:      data.Image,
			Rating:     data.Rating,
			JudulFoto:  data.JudulFoto,
			NamaLokasi: data.NamaLokasi,
			Deskripsi:  data.Deskripsi,
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

func UpdateGalery(id uuid.UUID, req *payload.UpdateGalery) (err error) {
	if _, err := database.GetGaleryById(id); err != nil {
		return err
	}
	gal := models.TbGalery{
		Image: req.Image,
	}
	if err := database.UpdateGalery(id, &gal); err != nil {
		return err
	}
	return nil

}

func DeleteGalery(id uuid.UUID) (err error) {
	if _, err := database.GetGaleryById(id); err != nil {
		return err
	}
	err = database.DeleteGalery(id)
	if err != nil {
		fmt.Println("Delete: error deleting galery, err:", err)
		return err
	}
	return nil
}
