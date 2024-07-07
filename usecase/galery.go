package usecase

import (
	"demonstrasi/config"
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

func GetGalery() ([]payload.GetGaleryRespone, error) {
	var galeries []struct {
		ID            string  `gorm:"column:id"`
		IDInformation string  `gorm:"column:id_information"`
		Image         string  `gorm:"column:image"`
		Rating        float64 `gorm:"column:rating"`
		JudulFoto     string  `gorm:"column:judul_foto"`
		NamaLokasi    string  `gorm:"column:nama_lokasi"`
		Deskripsi     string  `gorm:"column:deskripsi"`
	}

	query := `
       SELECT g.id, g.id_information, g.image, COALESCE(AVG(c.rating), 0) as rating,
               i.id_login, i.judul_foto, i.nama_lokasi, i.deskripsi
        FROM tb_galeries g
        LEFT JOIN tb_comments c ON g.id = c.id_galery
        LEFT JOIN tb_informations i ON g.id_information = i.id
        GROUP BY g.id, g.id_information, i.id_login, i.judul_foto, i.nama_lokasi, i.deskripsi
    `
	err := config.DB.Raw(query).Scan(&galeries).Error
	if err != nil {
		return nil, fmt.Errorf("error querying galery: %v", err)
	}

	var responses []payload.GetGaleryRespone
	for _, galery := range galeries {
		info := payload.GetInformationForGallery{
			ID:         galery.ID,
			JudulFoto:  galery.JudulFoto,
			NamaLokasi: galery.NamaLokasi,
			Deskripsi:  galery.Deskripsi,
		}

		response := payload.GetGaleryRespone{
			Image:                   galery.Image,
			Rating:                  galery.Rating,
			GetInformationForGalery: info,
		}
		responses = append(responses, response)
	}

	return responses, nil
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
