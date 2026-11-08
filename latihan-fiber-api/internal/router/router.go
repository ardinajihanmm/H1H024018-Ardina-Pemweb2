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

	// Seluruh resource dilindungi API Key.
	protected := api.Group(
		"",
		appmiddleware.KunciApi(),
	)

	// CRUD MAHASISWA
	mahasiswa := protected.Group("/mahasiswa")

	mahasiswa.Get("/", handler.DaftarMahasiswa)
	mahasiswa.Post("/", handler.BuatMahasiswa)
	mahasiswa.Get("/nim/:nim", handler.CariMahasiswaByNIM)
	mahasiswa.Get("/:id", handler.DetailMahasiswa)
	mahasiswa.Put("/:id", handler.PerbaruiMahasiswa)
	mahasiswa.Delete("/:id", handler.HapusMahasiswa)
	mahasiswa.Patch("/:id/pulihkan", handler.PulihkanMahasiswa)

	// CRUD MATA KULIAH
	matakuliah := protected.Group("/matakuliah")

	matakuliah.Get("/", handler.DaftarMatakuliah)
	matakuliah.Post("/", handler.BuatMatakuliah)
	matakuliah.Get("/:id", handler.DetailMatakuliah)
	matakuliah.Put("/:id", handler.PerbaruiMatakuliah)
	matakuliah.Delete("/:id", handler.HapusMatakuliah)

	// PROGRAM STUDI
	protected.Get(
		"/program-studi",
		handler.DaftarProgramStudi,
	)
}
