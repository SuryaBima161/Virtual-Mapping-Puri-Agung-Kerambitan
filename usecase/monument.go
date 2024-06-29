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

func UpdateMonument(id uuid.UUID, req *payload.UpdateMonument) (err error) {
	if _, err := database.GetMonumentById(id); err != nil {
		return err
	}
	mon := models.TbMonument{
		Image: req.Image,
	}
	if err := database.UpdateMonument(id, &mon); err != nil {
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
