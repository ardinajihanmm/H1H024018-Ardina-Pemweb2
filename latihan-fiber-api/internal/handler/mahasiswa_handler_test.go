package handler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"

	"latihan-fiber-api/internal/handler"
	"latihan-fiber-api/internal/router"
)

// Gunakan router produksi untuk memeriksa endpoint publik /api/status.
func TestStatusEndpoint(t *testing.T) {
	app := fiber.New()
	router.Daftarkan(app)
	permintaan := httptest.NewRequest(http.MethodGet, "/api/status", nil)
	respons, err := app.Test(permintaan)
	if err != nil {
		t.Fatal(err)
	}
	defer respons.Body.Close()
	if respons.StatusCode != http.StatusOK {
		t.Fatalf("diharapkan 200, diterima %d", respons.StatusCode)
	}
	var isi struct {
		Sukses bool `json:"sukses"`
	}
	if err := json.NewDecoder(respons.Body).Decode(&isi); err != nil {
		t.Fatal(err)
	}
	if !isi.Sukses {
		t.Fatal("endpoint status seharusnya sukses")
	}
}

func TestProfilTanpaTokenDitolak(t *testing.T) {
	app := fiber.New()
	router.Daftarkan(app)
	respons, err := app.Test(httptest.NewRequest(http.MethodGet, "/api/auth/profil", nil))
	if err != nil {
		t.Fatal(err)
	}
	defer respons.Body.Close()
	if respons.StatusCode != http.StatusUnauthorized {
		t.Fatalf("diharapkan 401, diterima %d", respons.StatusCode)
	}
}

func TestDetailMahasiswaIdBukanAngka(t *testing.T) {
	app := fiber.New()
	app.Get("/api/mahasiswa/:id", handler.DetailMahasiswa)
	respons, err := app.Test(httptest.NewRequest(http.MethodGet, "/api/mahasiswa/abc", nil))
	if err != nil {
		t.Fatal(err)
	}
	defer respons.Body.Close()
	if respons.StatusCode != http.StatusBadRequest {
		t.Fatalf("diharapkan 400, diterima %d", respons.StatusCode)
	}
}

func TestBuatMahasiswaNimTidakValid(t *testing.T) {
	// Validasi terjadi sebelum akses database; test ini tidak memerlukan MySQL.
	app := fiber.New()
	app.Post("/api/mahasiswa", handler.BuatMahasiswa)
	body := `{"program_studi_id":1,"nim":"A12345678","nama":"Mahasiswa Uji","email":"test.modul10@example.com","angkatan":2025,"ipk":3.5,"aktif":true}`
	permintaan := httptest.NewRequest(http.MethodPost, "/api/mahasiswa", strings.NewReader(body))
	permintaan.Header.Set("Content-Type", "application/json")
	respons, err := app.Test(permintaan)
	if err != nil {
		t.Fatal(err)
	}
	defer respons.Body.Close()
	if respons.StatusCode != http.StatusUnprocessableEntity {
		teks, _ := io.ReadAll(respons.Body)
		t.Fatalf("diharapkan 422, diterima %d: %s", respons.StatusCode, teks)
	}
}
