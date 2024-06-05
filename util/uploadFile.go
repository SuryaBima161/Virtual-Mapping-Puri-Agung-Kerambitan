package util

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func UploadFile(file *multipart.FileHeader) (string, error) {
	// Buka file tersebut
	src, err := file.Open()
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer src.Close()

	// Tentukan lokasi tujuan penyimpanan file di lokal
	dstPath := filepath.Join("uploads", file.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer dst.Close()

	// Salin konten file ke lokasi tujuan
	if _, err = io.Copy(dst, src); err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return dstPath, nil
}
