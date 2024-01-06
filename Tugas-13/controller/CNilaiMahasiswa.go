package controller

import (
	"encoding/json"
	"net/http"
	"sync"
	Model "tugas-13/model"
)

func NilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		getNilaiMahasiswa(w, r)
	case "POST":
		postNilaiMahasiswa(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(map[string]string{
			"Message": "Method Not Allowed !",
		})
		w.Write([]byte(res))
	}
}

var (
	nilaiNilaiMahasiswa = []Model.NilaiMahasiswa{}
	lastID              uint
	idMutex             sync.Mutex
)

func getNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	// Return all nilaiNilaiMahasiswa as JSON
	if len(nilaiNilaiMahasiswa) == 0 {
		w.WriteHeader(404)
		res, _ := json.Marshal(map[string]string{
			"Message": "Data Is Empty !",
		})
		w.Write([]byte(res))
		return
	}
	res, _ := json.Marshal(map[string]any{
		"Nilai_mahasiswa": nilaiNilaiMahasiswa,
	})
	w.Write([]byte(res))
}
func postNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	// Parse JSON payload
	var nilaiMahasiswa Model.NilaiMahasiswa
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&nilaiMahasiswa)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	//membuat id
	idMutex.Lock()
	nilaiMahasiswa.ID = lastID + 1
	lastID++
	idMutex.Unlock()
	//validasi nilai
	if nilaiMahasiswa.Nilai > 100 {
		http.Error(w, "Nilai dari 0 sampai 100", http.StatusBadRequest)
		return
	}
	// Index nilai
	if nilaiMahasiswa.Nilai >= 80 {
		nilaiMahasiswa.IndeksNilai = "A"
	} else if nilaiMahasiswa.Nilai >= 70 {
		nilaiMahasiswa.IndeksNilai = "B"
	} else if nilaiMahasiswa.Nilai >= 60 {
		nilaiMahasiswa.IndeksNilai = "C"
	} else if nilaiMahasiswa.Nilai >= 50 {
		nilaiMahasiswa.IndeksNilai = "D"
	} else {
		nilaiMahasiswa.IndeksNilai = "E"
	}
	res, _ := json.Marshal(nilaiMahasiswa)
	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilaiMahasiswa)
	w.Write([]byte(res))
}
