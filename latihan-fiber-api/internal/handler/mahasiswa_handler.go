package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"
)

func idMahasiswa(c fiber.Ctx) (uint, error) {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil || id == 0 {
		return 0, fiber.NewError(
			fiber.StatusBadRequest,
			"ID harus berupa angka positif",
		)
	}

	return uint(id), nil
}

func DaftarMahasiswa(c fiber.Ctx) error {
	halaman, _ := strconv.Atoi(c.Query("halaman", "1"))
	perHalaman, _ := strconv.Atoi(c.Query("per_halaman", "10"))
	angkatan, _ := strconv.Atoi(c.Query("angkatan", "0"))

	if halaman < 1 {
		halaman = 1
	}

	if perHalaman < 1 || perHalaman > 100 {
		perHalaman = 10
	}

	filter := repository.FilterMahasiswa{
		Cari:       c.Query("cari", ""),
		Angkatan:   angkatan,
		Halaman:    halaman,
		PerHalaman: perHalaman,
	}

	daftar, total, err := repository.Ambil(filter)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal mengambil data",
		)
	}

	totalHalaman := int(total) / perHalaman

	if int(total)%perHalaman != 0 {
		totalHalaman++
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   daftar,
		"meta": fiber.Map{
			"halaman":       halaman,
			"per_halaman":   perHalaman,
			"total":         total,
			"total_halaman": totalHalaman,
		},
	})
}

func DetailMahasiswa(c fiber.Ctx) error {
	id, err := idMahasiswa(c)
	if err != nil {
		return err
	}

	mahasiswa, err := repository.AmbilPerId(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(
				fiber.StatusNotFound,
				"Mahasiswa tidak ditemukan",
			)
		}

		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal mengambil data",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   mahasiswa,
	})
}

func CariMahasiswaByNIM(c fiber.Ctx) error {
	nim := c.Params("nim")

	mahasiswa, err := repository.AmbilPerNIM(nim)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(
				fiber.StatusNotFound,
				"Mahasiswa dengan NIM tersebut tidak ditemukan",
			)
		}

		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal mengambil data",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   mahasiswa,
	})
}

func BuatMahasiswa(c fiber.Ctx) error {
	var permintaan PermintaanMahasiswa

	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"Format body tidak valid",
		)
	}

	if permintaan.ProgramStudiID == 0 ||
		permintaan.Nim == "" ||
		permintaan.Nama == "" ||
		permintaan.Email == "" ||
		permintaan.Angkatan <= 0 ||
		permintaan.Ipk < 0 ||
		permintaan.Ipk > 4 {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"Data mahasiswa tidak lengkap atau tidak valid",
		)
	}

	aktif := true
	if permintaan.Aktif != nil {
		aktif = *permintaan.Aktif
	}

	mahasiswa := model.Mahasiswa{
		ProgramStudiID: permintaan.ProgramStudiID,
		Nim:            permintaan.Nim,
		Nama:           permintaan.Nama,
		Email:          permintaan.Email,
		Angkatan:       permintaan.Angkatan,
		Ipk:            permintaan.Ipk,
		Aktif:          aktif,
	}

	if err := repository.Simpan(&mahasiswa); err != nil {
		return fiber.NewError(
			fiber.StatusUnprocessableEntity,
			"Gagal menyimpan data, periksa NIM, email, dan program studi",
		)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil dibuat",
		"data":   mahasiswa,
	})
}

func PerbaruiMahasiswa(c fiber.Ctx) error {
	id, err := idMahasiswa(c)
	if err != nil {
		return err
	}

	mahasiswa, err := repository.AmbilPerId(id)
	if err != nil {
		return fiber.NewError(
			fiber.StatusNotFound,
			"Mahasiswa tidak ditemukan",
		)
	}

	var permintaan PermintaanMahasiswa

	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"Format body tidak valid",
		)
	}

	if permintaan.Nama != "" {
		mahasiswa.Nama = permintaan.Nama
	}

	if permintaan.Nim != "" {
		mahasiswa.Nim = permintaan.Nim
	}

	if permintaan.Email != "" {
		mahasiswa.Email = permintaan.Email
	}

	if permintaan.Angkatan > 0 {
		mahasiswa.Angkatan = permintaan.Angkatan
	}

	if permintaan.Ipk > 0 {
		mahasiswa.Ipk = permintaan.Ipk
	}

	if permintaan.ProgramStudiID > 0 {
		mahasiswa.ProgramStudiID = permintaan.ProgramStudiID
	}

	if permintaan.Aktif != nil {
		mahasiswa.Aktif = *permintaan.Aktif
	}

	if err := repository.Perbarui(&mahasiswa); err != nil {
		return fiber.NewError(
			fiber.StatusUnprocessableEntity,
			"Gagal memperbarui data",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil diperbarui",
		"data":   mahasiswa,
	})
}

func HapusMahasiswa(c fiber.Ctx) error {
	id, err := idMahasiswa(c)
	if err != nil {
		return err
	}

	mahasiswa, err := repository.AmbilPerId(id)
	if err != nil {
		return fiber.NewError(
			fiber.StatusNotFound,
			"Mahasiswa tidak ditemukan",
		)
	}

	if err := repository.Hapus(&mahasiswa); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal menghapus data",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil dihapus (soft delete)",
	})
}

func PulihkanMahasiswa(c fiber.Ctx) error {
	id, err := idMahasiswa(c)
	if err != nil {
		return err
	}

	_, err = repository.AmbilTerhapus(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(
				fiber.StatusNotFound,
				"Data mahasiswa terhapus tidak ditemukan",
			)
		}

		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal mencari data terhapus",
		)
	}

	if err := repository.Pulihkan(id); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal memulihkan data",
		)
	}

	mahasiswa, err := repository.AmbilPerId(id)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Data dipulihkan tetapi gagal diambil",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mahasiswa berhasil dipulihkan",
		"data":   mahasiswa,
	})
}
