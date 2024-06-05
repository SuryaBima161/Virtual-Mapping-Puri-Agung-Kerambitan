package database

import (
	"demonstrasi/config"
	"demonstrasi/models"

	uuid "github.com/satori/go.uuid"
)

func CreateComment(req *models.TbComment) error {
	if err := config.DB.Save(&req).Error; err != nil {
		return err
	}
	return nil
}

func GetComment() ([]models.TbComment, error) {
	var comment []models.TbComment
	db := config.DB
	if err := db.Find(&comment).Error; err != nil {
		return comment, err
	}
	return comment, nil
}

func GetCommentById(id uuid.UUID) (resp models.TbComment, err error) {
	if err := config.DB.Where("id = ?", id).First(&resp).Error; err != nil {
		return resp, err
	}
	return
}

func DeleteComment(id uuid.UUID) error {
	inf := models.TbComment{}
	if err := config.DB.Model(&inf).Where("id = ?", id).Delete(&inf).Error; err != nil {
		return err
	}
	return nil
}
