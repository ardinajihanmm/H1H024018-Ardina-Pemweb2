package middleware

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"

	"latihan-fiber-api/internal/auth"
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
)

func Autentikasi() fiber.Handler {
	konf := config.Muat()

	return func(c fiber.Ctx) error {
		header := c.Get("Authorization")

		if header == "" {
			return fiber.NewError(
				401,
				"Header Authorization tidak ditemukan",
			)
		}

		bagian := strings.Fields(header)

		if len(bagian) != 2 || bagian[0] != "Bearer" {
			return fiber.NewError(
				401,
				"Format header Authorization tidak sesuai",
			)
		}

		klaim, err := auth.PeriksaToken(bagian[1], konf.JwtSecret)
		if err != nil {
			return fiber.NewError(
				401,
				"Token tidak valid atau sudah kedaluwarsa",
			)
		}

		id, err := strconv.ParseUint(klaim.Subject, 10, 64)
		if err != nil {
			return fiber.NewError(401, "Isi token tidak sesuai")
		}

		// Periksa token yang sudah logout.
		var jumlah int64
		err = config.DB.Model(&model.TokenDiblokir{}).
			Where("token_id = ?", klaim.ID).
			Count(&jumlah).Error

		if err != nil {
			return fiber.NewError(500, "Gagal memeriksa token")
		}

		if jumlah > 0 {
			return fiber.NewError(
				401,
				"Token sudah dicabut melalui logout",
			)
		}

		var pengguna model.Pengguna
		if err := config.DB.First(&pengguna, uint(id)).Error; err != nil {
			return fiber.NewError(401, "Pengguna tidak ditemukan")
		}
		c.Locals("id_pengguna", uint(id))
		c.Locals("peran", pengguna.Peran)
		c.Locals("token_id", klaim.ID)
		c.Locals("token_exp", klaim.ExpiresAt)

		return c.Next()
	}
}

func HanyaPeran(peranDiizinkan ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		peran, cocok := c.Locals("peran").(string)

		if !cocok {
			return fiber.NewError(
				403,
				"Peran pengguna tidak diketahui",
			)
		}

		for _, izin := range peranDiizinkan {
			if izin == peran {
				return c.Next()
			}
		}

		return fiber.NewError(
			403,
			"Anda tidak berwenang mengakses sumber daya ini",
		)
	}
}
