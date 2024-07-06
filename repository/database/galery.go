package database

import (
	"demonstrasi/config"
	"demonstrasi/models"
	"demonstrasi/models/payload"
	"fmt"

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

func GetGalery() ([]payload.GetGaleryRespone, error) {
	var galeries []payload.GetGaleryRespone
	query := `
       SELECT g.id as id, g.image, COALESCE(AVG(c.rating), 0) as rating
        FROM tb_galeries g
        LEFT JOIN tb_comments c ON g.id = c.id_galery
        GROUP BY g.id
    `
	err := config.DB.Raw(query).Scan(&galeries).Error
	if err != nil {
		return nil, fmt.Errorf("error querying galery: %v", err)
	}
	return galeries, nil
}

func GetGaleryWithAvgRating() ([]payload.GaleryWithRating, error) {
	var galeries []payload.GaleryWithRating
	query := `
       SELECT g.id as id, g.image, COALESCE(AVG(c.rating), 0) as rating
        FROM tb_galeries g
        LEFT JOIN tb_comments c ON g.id = c.id_galery
        GROUP BY g.id
        ORDER BY rating DESC
        LIMIT 4
    `
	err := config.DB.Raw(query).Scan(&galeries).Error
	if err != nil {
		return nil, fmt.Errorf("error querying galery: %v", err)
	}
	return galeries, nil
}
