package validasi

import "testing"

type dataNIM struct {
	Nim string `validate:"required,nim_h"`
}

func TestNimHarusDiawaliH(t *testing.T) {
	if Periksa(dataNIM{Nim: "A123"}) == nil {
		t.Fatal("NIM awalan A seharusnya ditolak")
	}
}

func TestNimHanyaHurufHDanAngka(t *testing.T) {
	if Periksa(dataNIM{Nim: "H12A3"}) == nil {
		t.Fatal("NIM berisi huruf selain H seharusnya ditolak")
	}
}

func TestNimValid(t *testing.T) {
	if galat := Periksa(dataNIM{Nim: "H12345678"}); galat != nil {
		t.Fatalf("NIM seharusnya valid: %+v", galat)
	}
}
