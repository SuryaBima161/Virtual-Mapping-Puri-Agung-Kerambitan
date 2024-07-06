package database

import (
	"demonstrasi/config"
	"demonstrasi/models"
	"fmt"

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

func GetMonumentById(id uuid.UUID) (models.TbMonument, error) {
	var monument models.TbMonument
	err := config.DB.
		Preload("TbInformation").
		Where("id = ?", id).
		First(&monument).
		Error
	if err != nil {
		return monument, fmt.Errorf("error querying monument: %v", err)
	}
	return monument, nil
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
	if err := db.Preload("TbInformation").Find(&monument).Error; err != nil {
		return monument, err
	}
	return monument, nil
}
