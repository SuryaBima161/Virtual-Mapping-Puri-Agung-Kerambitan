package database

import (
	"demonstrasi/config"
	"demonstrasi/models"

	uuid "github.com/satori/go.uuid"
)

func CreateInformation(req *models.TbInformation) error {
	if err := config.DB.Save(&req).Error; err != nil {
		return err
	}
	return nil
}

func UpdateInfortmation(id uuid.UUID, inf *models.TbInformation) error {
	if err := config.DB.Model(inf).Where("id = ?", id).Updates(inf).Error; err != nil {
		return err
	}
	return nil
}

func GetInformationById(id uuid.UUID) (resp models.TbInformation, err error) {
	if err := config.DB.Where("id = ?", id).First(&resp).Error; err != nil {
		return resp, err
	}
	return
}

func DeleteInfortmation(id uuid.UUID) error {
	inf := models.TbInformation{}
	if err := config.DB.Model(&inf).Where("id = ?", id).Delete(&inf).Error; err != nil {
		return err
	}
	return nil
}

func GetInformation() ([]models.TbInformation, error) {
	var inf []models.TbInformation
	db := config.DB
	if err := db.Find(&inf).Error; err != nil {
		return inf, err
	}
	return inf, nil
}
