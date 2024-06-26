package payload

import uuid "github.com/satori/go.uuid"

type AddInformation struct {
	JudulFoto  string `json:"judul_foto" form:"judul_foto"`
	NamaLokasi string `json:"nama_lokasi" form:"nama_lokasi"`
	Deskripsi  string `json:"deskripsi" form:"deskripsi"`
}

type UpdateInfortmation struct {
	JudulFoto  string `json:"judul_foto" form:"judul_foto"`
	NamaLokasi string `json:"nama_lokasi" form:"nama_lokasi"`
	Deskripsi  string `json:"deskripsi" form:"deskripsi"`
}

type GetInformationRespone struct {
	Id_Information uuid.UUID `json:"id_information" form:"id_information"`
	JudulFoto      string    `json:"judul_foto" form:"judul_foto"`
	NamaLokasi     string    `json:"nama_lokasi" form:"nama_lokasi"`
	Deskripsi      string    `json:"deskripsi" form:"deskripsi"`
}
