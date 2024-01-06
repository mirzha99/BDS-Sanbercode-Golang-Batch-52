package nilai

import (
	"encoding/json"
	"log"
	"net/http"
	"tugas-15/config"
	Model "tugas-15/model"

	"github.com/julienschmidt/httprouter"
)

func GetNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.MySQL()
	if err != nil {
		log.Fatal(err)
	}

	QueryString := "SELECT nilai.id, nilai.indeks, nilai.skor, mahasiswa.id AS MahasiswaID, mahasiswa.nama AS MahasiswaNama, mata_kuliah.id AS MatakuliahID, mata_kuliah.nama AS MatakuliahNama FROM nilai JOIN mahasiswa ON nilai.mahasiswa_id = mahasiswa.id JOIN mata_kuliah ON nilai.matakuliah_id = mata_kuliah.id"
	row, err := db.Query(QueryString)
	if err != nil {
		log.Fatal(err)
	}
	// Slice untuk menyimpan hasil join
	var hasilJoin []map[string]interface{}
	for row.Next() {

		var nilai Model.Nilai
		var mahasiswa Model.Mahasiswa
		var matakuliah Model.Matakuliah

		err := row.Scan(&nilai.ID, &nilai.Indeks, &nilai.Skor, &mahasiswa.ID, &mahasiswa.Nama, &matakuliah.ID, &matakuliah.Nama)
		if err != nil {
			log.Fatal(err)
		}
		hasilJoin = append(hasilJoin, map[string]interface{}{
			"nilai_id":        nilai.ID,
			"indeks":          nilai.Indeks,
			"skor":            nilai.Skor,
			"mahasiswa_id":    mahasiswa.ID,
			"mahasiswa_nama":  mahasiswa.Nama,
			"matakuliah_id":   matakuliah.ID,
			"matakuliah_nama": matakuliah.Nama,
		})
	}
	if len(hasilJoin) == 0 {
		response := map[string]string{
			"message": "Data Is Empty",
		}

		res, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
		return
	}
	response := map[string]interface{}{
		"nilai": hasilJoin,
	}

	res, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetNilaiById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.MySQL()
	if err != nil {
		log.Fatal(err)
	}
	id := ps.ByName("id")
	QueryString := "SELECT nilai.id, nilai.indeks, nilai.skor, mahasiswa.id AS MahasiswaID, mahasiswa.nama AS MahasiswaNama, mata_kuliah.id AS MatakuliahID, mata_kuliah.nama AS MatakuliahNama FROM nilai JOIN mahasiswa ON nilai.mahasiswa_id = mahasiswa.id JOIN mata_kuliah ON nilai.matakuliah_id = mata_kuliah.id WHERE nilai.id = ?"
	row, err := db.Query(QueryString, id)
	if err != nil {
		log.Fatal(err)
	}
	// Slice untuk menyimpan hasil join
	var hasilJoin []map[string]interface{}
	for row.Next() {

		var nilai Model.Nilai
		var mahasiswa Model.Mahasiswa
		var matakuliah Model.Matakuliah

		err := row.Scan(&nilai.ID, &nilai.Indeks, &nilai.Skor, &mahasiswa.ID, &mahasiswa.Nama, &matakuliah.ID, &matakuliah.Nama)
		if err != nil {
			log.Fatal(err)
		}
		hasilJoin = append(hasilJoin, map[string]interface{}{
			"nilai_id":        nilai.ID,
			"indeks":          nilai.Indeks,
			"skor":            nilai.Skor,
			"mahasiswa_id":    mahasiswa.ID,
			"mahasiswa_nama":  mahasiswa.Nama,
			"matakuliah_id":   matakuliah.ID,
			"matakuliah_nama": matakuliah.Nama,
		})
	}
	if len(hasilJoin) == 0 {
		response := map[string]string{
			"message": "Data nilai " + id + " Not Found",
		}

		res, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
		return
	}
	response := map[string]interface{}{
		"nilai": hasilJoin,
	}

	res, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func PostNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}
	var nilai Model.Nilai
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&nilai)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	//validasi id mahasiswa
	if !config.CheckIDdb("mahasiswa", nilai.Mahasiswa_Id) {
		res, _ := json.Marshal(map[string]any{
			"Message": "id mahasiswa tidak ada",
		})
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(res))
		return
	}
	//validasi id mata_kuliah
	if !config.CheckIDdb("mata_kuliah", nilai.Matakuliah_Id) {
		res, _ := json.Marshal(map[string]any{
			"Message": "id mata kuliah tidak ada",
		})
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(res))
		return
	}
	//set index nilai
	nilai.SetIndexNilai()
	rows, err := db.Exec("INSERT INTO `nilai` (`id`, `indeks`, `skor`, `mahasiswa_id`, `matakuliah_id`) VALUES (NULL, ?, ?, ?, ?)", nilai.Indeks, nilai.Skor, nilai.Mahasiswa_Id, nilai.Matakuliah_Id)
	if err != nil {
		log.Fatal(err)
	}

	row, _ := rows.RowsAffected()

	if row == 0 {
		res, _ := json.Marshal(map[string]any{
			"Message": "Created Failed",
		})
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(res))
		return
	}
	res, _ := json.Marshal(map[string]any{
		"Message": "Created Success",
	})
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write([]byte(res))
}
func PutNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}
	var nilai Model.Nilai
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&nilai)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	//validasi id mahasiswa
	if !config.CheckIDdb("mahasiswa", nilai.Mahasiswa_Id) {
		res, _ := json.Marshal(map[string]any{
			"Message": "id mahasiswa tidak ada",
		})
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(res))
		return
	}
	//validasi id mata_kuliah
	if !config.CheckIDdb("mata_kuliah", nilai.Matakuliah_Id) {
		res, _ := json.Marshal(map[string]any{
			"Message": "id mata kuliah tidak ada",
		})
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(res))
		return
	}
	//set index nilai
	nilai.SetIndexNilai()

	id := ps.ByName("id")
	rows, err := db.Exec("UPDATE `nilai` SET `indeks` = ?, `skor` = ?, `mahasiswa_id` = ?, `matakuliah_id` = ? WHERE `nilai`.`id` = ?", nilai.Indeks, nilai.Skor, nilai.Mahasiswa_Id, nilai.Matakuliah_Id, id)
	if err != nil {
		log.Fatal(err)
	}

	row, _ := rows.RowsAffected()

	if row == 0 {
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
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write([]byte(res))
}
func DeleteNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	//id
	id := ps.ByName("id")
	rows, err := db.Exec("DELETE FROM `nilai` WHERE `nilai`.`id` = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	row, _ := rows.RowsAffected()

	if row == 0 {
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
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write([]byte(res))
}
