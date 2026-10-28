package router

import (
	"github.com/gofiber/fiber/v3"

	"latihan-fiber-api/internal/handler"
	appmiddleware "latihan-fiber-api/internal/middleware"
)

func Daftarkan(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"sukses": true,
			"pesan":  "API Fiber aktif",
		})
	})

	mahasiswa := api.Group(
		"/mahasiswa",
		appmiddleware.KunciApi(),
	)

	mahasiswa.Get("/", handler.DaftarMahasiswa)
	mahasiswa.Get("/nim/:nim", handler.CariMahasiswaByNIM)
	mahasiswa.Get("/:id", handler.DetailMahasiswa)
}