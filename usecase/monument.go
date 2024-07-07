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

func CreateMonument(req *payload.MonumentRequest, image *multipart.FileHeader) error {
	result, err := util.UploadFileToCloudinary(image)
	if err != nil {
		return err
	}

	mon := models.TbMonument{
		Id_Information: req.InformationID,
		Image:          result,
	}

	if err := database.CreateMonument(&mon); err != nil {
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
			JudulFoto:  data.TbInformation.JudulFoto,
			NamaLokasi: data.TbInformation.NamaLokasi,
			Deskripsi:  data.TbInformation.Deskripsi,
			Image:      data.Image,
		}
	}

	return resp, nil
}

func GetMonumentById(id uuid.UUID) (payload.GetMonumentRespone, error) {
	monument, err := database.GetMonumentById(id)
	if err != nil {
		return payload.GetMonumentRespone{}, fmt.Errorf("failed to get monument: %v", err)
	}

	resp := payload.GetMonumentRespone{
		JudulFoto:  monument.TbInformation.JudulFoto,
		NamaLokasi: monument.TbInformation.NamaLokasi,
		Deskripsi:  monument.TbInformation.Deskripsi,
		Image:      monument.Image,
	}

	return resp, nil
}

func UpdateMonument(id uuid.UUID, req *payload.UpdateMonument, image *multipart.FileHeader) (err error) {
	if _, err := database.GetMonumentById(id); err != nil {
		return err
	}

	var imageUrl string
	if image != nil {
		result, err := util.UploadFileToCloudinary(image)
		if err != nil {
			return err
		}
		imageUrl = result
	} else {
		existingGalery, err := database.GetMonumentById(id)
		if err != nil {
			return err
		}
		imageUrl = existingGalery.Image
	}
	gal := models.TbMonument{
		Image: imageUrl,
	}
	if err := database.UpdateMonument(id, &gal); err != nil {
		return err
	}
	return nil
}

func DeleteMonument(id uuid.UUID) (err error) {
	if _, err := database.GetMonumentById(id); err != nil {
		return err
	}
	err = database.DeleteMonument(id)
	if err != nil {
		fmt.Println("Delete: error deleting monument, err:", err)
		return err
	}
	return nil
}
