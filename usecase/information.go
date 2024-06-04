package usecase

import (
	"demonstrasi/models"
	"demonstrasi/models/payload"
	"demonstrasi/repository/database"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func CreateInformation(req *payload.AddInformation) error {
	inf := models.TbInformation{
		NamaLokasi: req.NamaLokasi,
		JudulFoto:  req.JudulFoto,
		Deskripsi:  req.Deskripsi,
	}
	if err := database.CreateInformation(&inf); err != nil {
		return err
	}
	return nil
}

func UpdateInfortmation(id uuid.UUID, req *payload.UpdateInfortmation) (err error) {
	if _, err := database.GetInformationById(id); err != nil {
		return err
	}
	inf := models.TbInformation{
		NamaLokasi: req.NamaLokasi,
		JudulFoto:  req.JudulFoto,
		Deskripsi:  req.Deskripsi,
	}
	if err := database.UpdateInfortmation(id, &inf); err != nil {
		return err
	}
	return nil

}

func GetInformation() (resp []payload.GetInformationRespone, err error) {
	inf, err := database.GetInformation()
	if err != nil {
		return []payload.GetInformationRespone{}, err
	}
	resp = make([]payload.GetInformationRespone, len(inf))
	for i, data := range inf {
		resp[i] = payload.GetInformationRespone{
			JudulFoto:  data.JudulFoto,
			NamaLokasi: data.JudulFoto,
			Deskripsi:  data.Deskripsi,
		}
	}

	return resp, nil
}

func GetInformationById(id uuid.UUID) (resp payload.GetInformationRespone, err error) {
	inf, err := database.GetInformationById(id)
	if err != nil {
		return payload.GetInformationRespone{}, err
	}
	resp = payload.GetInformationRespone{
		NamaLokasi: inf.NamaLokasi,
		JudulFoto:  inf.JudulFoto,
		Deskripsi:  inf.Deskripsi,
	}
	return
}

func DeleteInfortmation(id uuid.UUID) (err error) {
	if _, err := database.GetInformationById(id); err != nil {
		return err
	}
	err = database.DeleteInfortmation(id)
	if err != nil {
		fmt.Println("Delete: error deleting information, err:", err)
		return err
	}
	return nil
}
