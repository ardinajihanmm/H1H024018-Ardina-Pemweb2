package handler

import (
	"sort"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"latihan-fiber-api/internal/model"
)

func DaftarMahasiswa(c fiber.Ctx) error {
	angkatan := c.Query("angkatan")
	urut := c.Query("urut")

	var hasil []model.Mahasiswa

	for _, mahasiswa := range model.Daftar {
		if angkatan == "" {
			hasil = append(hasil, mahasiswa)
			continue
		}

		tahun, err := strconv.Atoi(angkatan)
		if err != nil {
			return fiber.NewError(
				fiber.StatusBadRequest,
				"Parameter angkatan harus berupa angka",
			)
		}

		if mahasiswa.Angkatan == tahun {
			hasil = append(hasil, mahasiswa)
		}
	}

	switch urut {
	case "nama":
		sort.Slice(hasil, func(i, j int) bool {
			return hasil[i].Nama < hasil[j].Nama
		})
	case "ipk":
		sort.Slice(hasil, func(i, j int) bool {
			return hasil[i].Ipk > hasil[j].Ipk
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   hasil,
	})
}

func DetailMahasiswa(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"ID mahasiswa harus berupa angka",
		)
	}

	for _, mahasiswa := range model.Daftar {
		if mahasiswa.ID == id {
			return c.JSON(fiber.Map{
				"sukses": true,
				"data":   mahasiswa,
			})
		}
	}

	return fiber.NewError(
		fiber.StatusNotFound,
		"Mahasiswa tidak ditemukan",
	)
}

func CariMahasiswaByNIM(c fiber.Ctx) error {
	nim := c.Params("nim")

	for _, mahasiswa := range model.Daftar {
		if mahasiswa.Nim == nim {
			return c.JSON(fiber.Map{
				"sukses": true,
				"data":   mahasiswa,
			})
		}
	}

	return fiber.NewError(
		fiber.StatusNotFound,
		"Mahasiswa dengan NIM tersebut tidak ditemukan",
	)
}