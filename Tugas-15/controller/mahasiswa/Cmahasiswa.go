package mahasiswa

import (
	"encoding/json"
	"log"
	"net/http"
	"tugas-15/config"
	Model "tugas-15/model"

	"github.com/julienschmidt/httprouter"
)

func GetMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	var mahasiswas []Model.Mahasiswa
	rows, err := db.Query("SELECT * FROM mahasiswa")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var mahasiswa Model.Mahasiswa
		err := rows.Scan(
			&mahasiswa.ID,
			&mahasiswa.Nama,
		)

		if err != nil {
			log.Fatal(err)
		}
		mahasiswas = append(mahasiswas, mahasiswa)
	}
	if len(mahasiswas) == 0 {
		respose := map[string]string{
			"message": "Data Mahasiswa Kosong",
		}
		res, _ := json.Marshal(respose)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(res))
		return
	}
	respose := map[string]interface{}{
		"mahasiswa": mahasiswas,
	}
	res, _ := json.Marshal(respose)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
func PostMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}
	var mahasiswa Model.Mahasiswa
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&mahasiswa)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	rows, err := db.Exec("INSERT INTO `mahasiswa` (`id`, `nama`) VALUES (NULL, ?)", mahasiswa.Nama)
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
func PutMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}
	var mahasiswa Model.Mahasiswa
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&mahasiswa)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	//id
	id := ps.ByName("id")
	rows, err := db.Exec("UPDATE `mahasiswa` SET `nama` = ? WHERE `mahasiswa`.`id` = ?", mahasiswa.Nama, id)
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
func DeleteMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	//id
	id := ps.ByName("id")
	rows, err := db.Exec("DELETE FROM `mahasiswa` WHERE `mahasiswa`.`id` = ?", id)
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
