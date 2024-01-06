package Model

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
	Created_at, Updated_at        string
}

func (nm *NilaiMahasiswa) SetIndexNilai() {
	if nm.Nilai >= 80 {
		nm.IndeksNilai = "A"
	} else if nm.Nilai >= 70 {
		nm.IndeksNilai = "B"
	} else if nm.Nilai >= 60 {
		nm.IndeksNilai = "C"
	} else if nm.Nilai >= 50 {
		nm.IndeksNilai = "D"
	} else {
		nm.IndeksNilai = "E"
	}
}
