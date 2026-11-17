package validasi

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"regexp"
)

var pemeriksa = validator.New()
var polaNIM = regexp.MustCompile(`^H[0-9]+$`)

func init() {
	if err := pemeriksa.RegisterValidation("nim_h", func(fl validator.FieldLevel) bool {
		return polaNIM.MatchString(fl.Field().String())
	}); err != nil {
		panic(err)
	}
}

type PesanGalat struct {
	Field string `json:"field"`
	Pesan string `json:"pesan"`
}

func Periksa(data any) []PesanGalat {
	err := pemeriksa.Struct(data)
	if err == nil {
		return nil
	}
	var kesalahan validator.ValidationErrors
	if !errors.As(err, &kesalahan) {
		return []PesanGalat{{Field: "data", Pesan: "Data tidak valid"}}
	}
	daftar := make([]PesanGalat, 0, len(kesalahan))
	for _, item := range kesalahan {
		daftar = append(daftar, PesanGalat{Field: item.Field(), Pesan: susunPesan(item)})
	}
	return daftar
}
func susunPesan(item validator.FieldError) string {
	switch item.Tag() {
	case "required":
		return "Field ini wajib diisi"
	case "email":
		return "Format email tidak valid"
	case "min":
		return "Nilai minimal adalah " + item.Param()
	case "max":
		return "Nilai maksimal adalah " + item.Param()
	case "gte":
		return "Nilai tidak boleh kurang dari " + item.Param()
	case "lte":
		return "Nilai tidak boleh lebih dari " + item.Param()
	case "nim_h":
		return "NIM harus diawali H dan diikuti angka"
	default:
		return "Nilai tidak memenuhi aturan " + item.Tag()
	}
}
