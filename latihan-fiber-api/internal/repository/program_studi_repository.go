package repository

import (
	"latihan-fiber-api/internal/config"
)

type ProgramStudiDenganJumlah struct {
	ID              uint   `json:"id"`
	Kode            string `json:"kode"`
	Nama            string `json:"nama"`
	Jenjang         string `json:"jenjang"`
	JumlahMahasiswa int64  `json:"jumlah_mahasiswa"`
}

func AmbilProgramStudi() ([]ProgramStudiDenganJumlah, error) {
	daftar := []ProgramStudiDenganJumlah{}

	err := config.DB.
		Table("program_studis AS ps").
		Select(
			"ps.id, ps.kode, ps.nama, ps.jenjang, " +
				"COUNT(m.id) AS jumlah_mahasiswa",
		).
		Joins(
			"LEFT JOIN mahasiswas AS m ON " +
				"m.program_studi_id = ps.id " +
				"AND m.deleted_at IS NULL",
		).
		Group("ps.id, ps.kode, ps.nama, ps.jenjang").
		Order("ps.nama ASC").
		Scan(&daftar).Error

	return daftar, err
}
