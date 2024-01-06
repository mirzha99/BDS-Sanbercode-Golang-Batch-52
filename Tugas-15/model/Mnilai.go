package Model

type Nilai struct {
	ID            int    `json:"id"`
	Indeks        string `json:"indeks"`
	Skor          uint   `json:"skor"`
	Mahasiswa_Id  int    `json:"mahasiswa_id"`
	Matakuliah_Id int    `json:"matakuliah_id"`
}

func (nm *Nilai) SetIndexNilai() {
	if nm.Skor >= 80 {
		nm.Indeks = "A"
	} else if nm.Skor >= 70 {
		nm.Indeks = "B"
	} else if nm.Skor >= 60 {
		nm.Indeks = "C"
	} else if nm.Skor >= 50 {
		nm.Indeks = "D"
	} else {
		nm.Indeks = "E"
	}
}
