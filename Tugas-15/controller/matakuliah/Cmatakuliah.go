package matakuliah

import (
	"encoding/json"
	"log"
	"net/http"
	"tugas-15/config"
	Model "tugas-15/model"

	"github.com/julienschmidt/httprouter"
)

func GetMatakuliah(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	var Matakuliahs []Model.Matakuliah
	rows, err := db.Query("SELECT * FROM mata_kuliah")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var Matakuliah Model.Matakuliah
		err := rows.Scan(
			&Matakuliah.ID,
			&Matakuliah.Nama,
		)

		if err != nil {
			log.Fatal(err)
		}
		Matakuliahs = append(Matakuliahs, Matakuliah)
	}
	if len(Matakuliahs) == 0 {
		respose := map[string]string{
			"message": "Data Mata kuliah Kosong",
		}
		res, _ := json.Marshal(respose)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(res))
		return
	}
	respose := map[string]interface{}{
		"Matakuliah": Matakuliahs,
	}
	res, _ := json.Marshal(respose)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
func PostMatakuliah(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}
	var Matakuliah Model.Matakuliah
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&Matakuliah)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	rows, err := db.Exec("INSERT INTO `mata_kuliah` (`id`, `nama`) VALUES (NULL, ?)", Matakuliah.Nama)
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
func PutMatakuliah(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}
	var Matakuliah Model.Matakuliah
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&Matakuliah)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	//id
	id := ps.ByName("id")
	rows, err := db.Exec("UPDATE `mata_kuliah` SET `nama` = ? WHERE `mata_kuliah`.`id` = ?", Matakuliah.Nama, id)
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
func DeleteMatakuliah(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	//id
	id := ps.ByName("id")
	rows, err := db.Exec("DELETE FROM `mata_kuliah` WHERE `mata_kuliah`.`id` = ?", id)
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
