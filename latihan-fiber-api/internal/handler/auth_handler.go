package handler

import (
	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"

	"latihan-fiber-api/internal/auth"
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/validasi"
)

func Registrasi(c fiber.Ctx) error {
	var permintaan PermintaanRegistrasi

	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(400, "Format body tidak valid")
	}

	if galat := validasi.Periksa(permintaan); galat != nil {
		return c.Status(422).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data yang dikirim tidak valid",
			"galat":  galat,
		})
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(permintaan.KataSandi),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return fiber.NewError(500, "Gagal memproses kata sandi")
	}

	pengguna := model.Pengguna{
		Nama:      permintaan.Nama,
		Email:     permintaan.Email,
		KataSandi: string(hash),
		Peran:     "mahasiswa",
	}

	if err := config.DB.Create(&pengguna).Error; err != nil {
		return fiber.NewError(422, "Gagal mendaftarkan pengguna; periksa apakah email sudah terdaftar")
	}

	return c.Status(201).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Registrasi berhasil",
		"data":   pengguna,
	})
}

func Login(c fiber.Ctx) error {
	var permintaan PermintaanLogin

	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(400, "Format body tidak valid")
	}

	if galat := validasi.Periksa(permintaan); galat != nil {
		return c.Status(422).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data yang dikirim tidak valid",
			"galat":  galat,
		})
	}

	var pengguna model.Pengguna

	err := config.DB.Where(
		"email = ?", permintaan.Email,
	).First(&pengguna).Error

	if err != nil {
		return fiber.NewError(401, "Email atau kata sandi tidak sesuai")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(pengguna.KataSandi),
		[]byte(permintaan.KataSandi),
	)
	if err != nil {
		return fiber.NewError(401, "Email atau kata sandi tidak sesuai")
	}

	konf := config.Muat()

	token, err := auth.TerbitkanToken(
		pengguna.ID,
		pengguna.Peran,
		konf.JwtSecret,
		konf.JwtTtlMenit,
	)
	if err != nil {
		return fiber.NewError(500, "Gagal menerbitkan token")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Login berhasil",
		"data": fiber.Map{
			"pengguna": pengguna,
			"token":    token,
			"tipe":     "Bearer",
		},
	})
}

func Profil(c fiber.Ctx) error {
	idPengguna := c.Locals("id_pengguna")

	var pengguna model.Pengguna

	if err := config.DB.First(&pengguna, idPengguna).Error; err != nil {
		return fiber.NewError(404, "Pengguna tidak ditemukan")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   pengguna,
	})
}
