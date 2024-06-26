package database

import (
	"demonstrasi/config"
	"demonstrasi/models"

	uuid "github.com/satori/go.uuid"
)

func CreateMonument(req *models.TbMonument) error {
	if err := config.DB.Save(&req).Error; err != nil {
		return err
	}
	return nil
}

func UpdateMonument(id uuid.UUID, monument *models.TbMonument) error {
	if err := config.DB.Model(monument).Where("id = ?", id).Updates(monument).Error; err != nil {
		return err
	}
	return nil
}

func GetMonumentById(id uuid.UUID) (resp models.TbMonument, err error) {
	if err := config.DB.Where("id = ?", id).First(&resp).Error; err != nil {
		return resp, err
	}
	return
}

func DeleteMonument(id uuid.UUID) error {
	monument := models.TbMonument{}
	if err := config.DB.Model(&monument).Where("id = ?", id).Delete(&monument).Error; err != nil {
		return err
	}
	return nil
}

func GetMonument() ([]models.TbMonument, error) {
	var monument []models.TbMonument
	db := config.DB
	if err := db.Find(&monument).Error; err != nil {
		return monument, err
	}
	return monument, nil
}
