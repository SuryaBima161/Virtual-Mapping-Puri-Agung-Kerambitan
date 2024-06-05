package database

import (
	"demonstrasi/config"
	"demonstrasi/models"

	uuid "github.com/satori/go.uuid"
)

func CreateGalery(req *models.TbGalery) error {
	if err := config.DB.Save(&req).Error; err != nil {
		return err
	}
	return nil
}

func UpdateGalery(id uuid.UUID, inf *models.TbGalery) error {
	if err := config.DB.Model(inf).Where("id = ?", id).Updates(inf).Error; err != nil {
		return err
	}
	return nil
}

func GetGaleryById(id uuid.UUID) (resp models.TbGalery, err error) {
	if err := config.DB.Where("id = ?", id).First(&resp).Error; err != nil {
		return resp, err
	}
	return
}

func DeleteGalery(id uuid.UUID) error {
	galery := models.TbGalery{}
	if err := config.DB.Model(&galery).Where("id = ?", id).Delete(&galery).Error; err != nil {
		return err
	}
	return nil
}

func GetGalery() ([]models.TbGalery, error) {
	var galery []models.TbGalery
	db := config.DB
	if err := db.Find(&galery).Error; err != nil {
		return galery, err
	}
	return galery, nil
}
