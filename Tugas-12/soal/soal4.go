package soal

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

func Soal4(w http.ResponseWriter, r *http.Request) {
	// Menggunakan data yang diberikan
	jariJari := 7.0
	tinggi := 10.0

	// Menghitung volume tabung
	volume := math.Pi * math.Pow(jariJari, 2) * tinggi

	// Menghitung luas alas tabung
	luasAlas := math.Pi * math.Pow(jariJari, 2)

	// Menghitung keliling alas tabung
	kelilingAlas := 2 * math.Pi * jariJari
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	hasil := fmt.Sprintf("jariJari : %v, tinggi: %v, volume : %v, luas alas: %v, keliling alas: %v", jariJari, tinggi, volume, luasAlas, kelilingAlas)
	response := map[string]string{
		"Hasil": hasil,
	}
	jsonResponse, _ := json.Marshal(&response)
	w.Write(jsonResponse)

}
