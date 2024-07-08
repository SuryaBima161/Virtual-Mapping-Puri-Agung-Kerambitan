package payload

import uuid "github.com/satori/go.uuid"

type AddGalery struct {
	InformationID uuid.UUID `json:"information_id" form:"information_id"`
	Image         string    `json:"image" form:"image"`
}

type UpdateGalery struct {
	Image string `json:"image" form:"image"`
}

type GetGaleryRespone struct {
	Id_galery               string                   `json:"id_galery" form:"id_galery"`
	Image                   string                   `json:"image" form:"image"`
	Rating                  float64                  `json:"rating" form:"rating"`
	GetInformationForGalery GetInformationForGallery `json:"information" form:"information"`
}

type GetInformationForGallery struct {
	ID         string `json:"id_information" form:"id_information"`
	JudulFoto  string `json:"judul_foto" form:"judul_foto"`
	NamaLokasi string `json:"nama_lokasi" form:"nama_lokasi"`
	Deskripsi  string `json:"deskripsi" form:"deskripsi"`
}
type GaleryWithRating struct {
	ID     uuid.UUID `json:"id"`
	Image  string    `json:"image"`
	Rating float64   `json:"rating"`
}
type GetHomePageRespone struct {
	ID     uuid.UUID `json:"id" form:"id"`
	Image  string    `json:"image" form:"image"`
	Rating float64   `json:"rating" form:"rating"`
}
