package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"
)

type PermintaanMatakuliah struct {
	Kode     string `json:"kode"`
	Nama     string `json:"nama"`
	Sks      int    `json:"sks"`
	Semester int    `json:"semester"`
}

func idMatakuliah(c fiber.Ctx) (uint, error) {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil || id == 0 {
		return 0, fiber.NewError(
			fiber.StatusBadRequest,
			"ID mata kuliah tidak valid",
		)
	}

	return uint(id), nil
}

func DaftarMatakuliah(c fiber.Ctx) error {
	daftar, err := repository.AmbilMatakuliah()

	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal mengambil data mata kuliah",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   daftar,
	})
}

func DetailMatakuliah(c fiber.Ctx) error {
	id, err := idMatakuliah(c)
	if err != nil {
		return err
	}

	matakuliah, err := repository.AmbilMatakuliahPerId(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(
			fiber.StatusNotFound,
			"Mata kuliah tidak ditemukan",
		)
	}

	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal mengambil mata kuliah",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   matakuliah,
	})
}

func BuatMatakuliah(c fiber.Ctx) error {
	var permintaan PermintaanMatakuliah

	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"Format body tidak valid",
		)
	}

	if permintaan.Kode == "" ||
		permintaan.Nama == "" ||
		permintaan.Sks <= 0 ||
		permintaan.Semester <= 0 {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"Data mata kuliah tidak lengkap",
		)
	}

	matakuliah := model.Matakuliah{
		Kode:     permintaan.Kode,
		Nama:     permintaan.Nama,
		Sks:      permintaan.Sks,
		Semester: permintaan.Semester,
	}

	if err := repository.SimpanMatakuliah(&matakuliah); err != nil {
		return fiber.NewError(
			fiber.StatusUnprocessableEntity,
			"Gagal menyimpan mata kuliah, periksa kode",
		)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mata kuliah berhasil ditambahkan",
		"data":   matakuliah,
	})
}

func PerbaruiMatakuliah(c fiber.Ctx) error {
	id, err := idMatakuliah(c)
	if err != nil {
		return err
	}

	matakuliah, err := repository.AmbilMatakuliahPerId(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(
				fiber.StatusNotFound,
				"Mata kuliah tidak ditemukan",
			)
		}
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal mengambil mata kuliah",
		)
	}

	var permintaan PermintaanMatakuliah

	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"Format body tidak valid",
		)
	}

	if permintaan.Kode != "" {
		matakuliah.Kode = permintaan.Kode
	}

	if permintaan.Nama != "" {
		matakuliah.Nama = permintaan.Nama
	}

	if permintaan.Sks > 0 {
		matakuliah.Sks = permintaan.Sks
	}

	if permintaan.Semester > 0 {
		matakuliah.Semester = permintaan.Semester
	}

	if err := repository.PerbaruiMatakuliah(&matakuliah); err != nil {
		return fiber.NewError(
			fiber.StatusUnprocessableEntity,
			"Gagal memperbarui mata kuliah",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mata kuliah berhasil diperbarui",
		"data":   matakuliah,
	})
}

func HapusMatakuliah(c fiber.Ctx) error {
	id, err := idMatakuliah(c)
	if err != nil {
		return err
	}

	matakuliah, err := repository.AmbilMatakuliahPerId(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(
				fiber.StatusNotFound,
				"Mata kuliah tidak ditemukan",
			)
		}
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal mengambil mata kuliah",
		)
	}

	if err := repository.HapusMatakuliah(&matakuliah); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal menghapus mata kuliah",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mata kuliah berhasil dihapus",
	})
}
