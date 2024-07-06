package payload

import uuid "github.com/satori/go.uuid"

type MonumentRequest struct {
	InformationID uuid.UUID `json:"information_id" form:"information_id"`
	Image         string    `json:"image" form:"image"`
}

type UpdateMonument struct {
	Image string `json:"image" form:"image"`
}
type GetMonumentRespone struct {
	JudulFoto  string `json:"judul_foto" form:"judul_foto"`
	NamaLokasi string `json:"nama_lokasi" form:"nama_lokasi"`
	Deskripsi  string `json:"deskripsi" form:"deskripsi"`
	Image      string `json:"image" form:"image"`
}
