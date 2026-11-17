package handler

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"

	"latihan-fiber-api/internal/auth"
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
)

func SegarkanToken(c fiber.Ctx) error {
	idPengguna, cocok := c.Locals("id_pengguna").(uint)
	if !cocok {
		return fiber.NewError(401, "Identitas pengguna tidak valid")
	}

	var pengguna model.Pengguna

	if err := config.DB.First(&pengguna, idPengguna).Error; err != nil {
		return fiber.NewError(401, "Pengguna tidak ditemukan")
	}

	konf := config.Muat()

	token, err := auth.TerbitkanToken(
		pengguna.ID,
		pengguna.Peran,
		konf.JwtSecret,
		konf.JwtTtlMenit,
	)
	if err != nil {
		return fiber.NewError(500, "Gagal memperbarui token")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Token berhasil diperbarui",
		"data": fiber.Map{
			"token": token,
			"tipe":  "Bearer",
		},
	})
}

func Logout(c fiber.Ctx) error {
	tokenID, cocok := c.Locals("token_id").(string)
	if !cocok || tokenID == "" {
		return fiber.NewError(401, "Identitas token tidak valid")
	}

	expiredAt, cocok := c.Locals("token_exp").(*jwt.NumericDate)
	if !cocok || expiredAt == nil {
		return fiber.NewError(401, "Waktu kedaluwarsa token tidak valid")
	}

	data := model.TokenDiblokir{
		TokenID:   tokenID,
		ExpiredAt: expiredAt.Time,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		return fiber.NewError(500, "Gagal mencabut token")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Logout berhasil",
		"waktu":  time.Now(),
	})
}
