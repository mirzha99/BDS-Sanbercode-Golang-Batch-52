package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"tugas-14/config"
	Model "tugas-14/model"

	"github.com/julienschmidt/httprouter"
)

const (
	table = "nilaimahasiswa"
)

func GetNilaiMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var nilaiNilaiMahasiswaList = []Model.NilaiMahasiswa{}
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC", table)
	rows, err := db.Query(queryText)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var nilaiMahasiswa = Model.NilaiMahasiswa{}

		err := rows.Scan(
			&nilaiMahasiswa.ID,
			&nilaiMahasiswa.Nama,
			&nilaiMahasiswa.MataKuliah,
			&nilaiMahasiswa.Nilai,
			&nilaiMahasiswa.IndeksNilai,
			&nilaiMahasiswa.Created_at,
			&nilaiMahasiswa.Updated_at)

		if err != nil {
			log.Fatal(err)
		}
		created, _ := strconv.Atoi(nilaiMahasiswa.Created_at)
		updated_at, _ := strconv.Atoi(nilaiMahasiswa.Updated_at)
		nilaiMahasiswa.Created_at = config.GetHumanTime(int64(created))
		nilaiMahasiswa.Updated_at = config.GetHumanTime(int64(updated_at))
		nilaiNilaiMahasiswaList = append(nilaiNilaiMahasiswaList, nilaiMahasiswa)
	}
	//check data
	if len(nilaiNilaiMahasiswaList) == 0 {
		w.WriteHeader(http.StatusNotFound)
		res, _ := json.Marshal(map[string]any{
			"message": "Data is Empty !",
		})
		w.Write(res)
		return
	}
	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	res, _ := json.Marshal(map[string]any{
		"Nilai_mahasiswa": nilaiNilaiMahasiswaList,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
func PostNilaiMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	// Parse JSON payload
	var nilaiMahasiswa Model.NilaiMahasiswa
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&nilaiMahasiswa)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	//db
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	//validasi nilai
	if nilaiMahasiswa.Nilai > 100 {
		http.Error(w, "Nilai dari 0 sampai 100", http.StatusBadRequest)
		return
	}
	// Index nilai
	nilaiMahasiswa.SetIndexNilai()
	//set Time
	nilaiMahasiswa.Created_at = strconv.Itoa(int(time.Now().Unix()))
	nilaiMahasiswa.Updated_at = strconv.Itoa(int(time.Now().Unix()))
	result, err := db.Exec("INSERT INTO `nilaimahasiswa` (`id`, `nama`, `matakuliah`, `nilai`, `indeksnilai`, `created_at`, `updated_at`) VALUES (NULL, ?, ?, ?, ?,?, ?)",
		nilaiMahasiswa.Nama, nilaiMahasiswa.MataKuliah, nilaiMahasiswa.Nilai, nilaiMahasiswa.IndeksNilai, nilaiMahasiswa.Created_at, nilaiMahasiswa.Updated_at)
	if err != nil {
		log.Fatal(err)
	}

	rowAff, _ := result.RowsAffected()

	if rowAff == 0 {
		res, _ := json.Marshal(map[string]any{
			"Message":        "Created Failed",
			"NilaiMahasiswa": nilaiMahasiswa,
		})
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(res))
		return
	}
	res, _ := json.Marshal(map[string]any{
		"Message":        "Created Success",
		"NilaiMahasiswa": nilaiMahasiswa,
	})
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(res))
}

func PutNilaiMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	w.Header().Set("Content-Type", "application/json")
	// Parse JSON payload
	var nilaiMahasiswa Model.NilaiMahasiswa
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&nilaiMahasiswa)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	//db
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	//validasi nilai
	if nilaiMahasiswa.Nilai > 100 {
		http.Error(w, "Nilai dari 0 sampai 100", http.StatusBadRequest)
		return
	}
	// Index nilai
	nilaiMahasiswa.SetIndexNilai()

	//set date
	nilaiMahasiswa.Updated_at = strconv.Itoa(int(time.Now().Unix()))
	result, err := db.Exec("UPDATE `nilaimahasiswa` SET `nama` = ?, `matakuliah` = ?, `nilai` = ?, `indeksnilai` = ?, `updated_at` = ? WHERE `nilaimahasiswa`.`id` = ?",
		nilaiMahasiswa.Nama, nilaiMahasiswa.MataKuliah, nilaiMahasiswa.Nilai, nilaiMahasiswa.IndeksNilai, nilaiMahasiswa.Updated_at, id)
	if err != nil {
		log.Fatal(err)
	}

	rowAff, _ := result.RowsAffected()

	if rowAff == 0 {
		res, _ := json.Marshal(map[string]any{
			"Message": "Update Failed",
		})
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(res))
		return
	}
	res, _ := json.Marshal(map[string]any{
		"Message": "Update Success",
	})
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(res))
}
func DeleteNilaiMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	//db
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}
	id := ps.ByName("id")
	result, err := db.Exec("DELETE FROM `nilaimahasiswa` WHERE `nilaimahasiswa`.`id` = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	rowAff, _ := result.RowsAffected()

	if rowAff == 0 {
		res, _ := json.Marshal(map[string]any{
			"Message": "Delete Failed",
		})
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(res))
		return
	}
	res, _ := json.Marshal(map[string]any{
		"Message": "Delete Success",
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
