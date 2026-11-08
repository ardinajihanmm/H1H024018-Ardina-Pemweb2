package repository

import (
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
)

func AmbilMatakuliah() ([]model.Matakuliah, error) {
	daftar := []model.Matakuliah{}

	err := config.DB.Order("nama ASC").
		Find(&daftar).Error

	return daftar, err
}

func AmbilMatakuliahPerId(id uint) (model.Matakuliah, error) {
	var matakuliah model.Matakuliah

	err := config.DB.First(&matakuliah, id).Error

	return matakuliah, err
}

func SimpanMatakuliah(matakuliah *model.Matakuliah) error {
	return config.DB.Create(matakuliah).Error
}

func PerbaruiMatakuliah(matakuliah *model.Matakuliah) error {
	return config.DB.Save(matakuliah).Error
}

func HapusMatakuliah(matakuliah *model.Matakuliah) error {
	return config.DB.Delete(matakuliah).Error
}
