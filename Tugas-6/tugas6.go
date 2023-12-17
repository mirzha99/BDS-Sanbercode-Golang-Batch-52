package main

import (
	"fmt"
	"math"
	"strings"
)

// func soal 1
func luaskelilingLingkaran(r, luasLigkaran, kelilingLingkaran *float64) {
	*luasLigkaran = math.Pi * (*r) * (*r)
	*kelilingLingkaran = 2 * math.Pi * (*r)
}

// func soal 2
func introduce(sentence *string, name, gender, jobs, age string) {
	if strings.ToLower(gender) == "laki-laki" || strings.ToLower(gender) == "cowok" {
		gender = "Pak"
	} else if strings.ToLower(gender) == "perempuan" || strings.ToLower(gender) == "cewek" {
		gender = "Bu"
	} else {
		gender = "false"
	}

	if gender == "false" {
		*sentence = "Tolak LGBT"
	} else {
		*sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", gender, name, jobs, age)
	}
}

// func soal 3
func tambahDataBuah(buah *[]string, namaBuah string) {
	*buah = append(*buah, namaBuah)
}

// func soal 4
func tambahDataFilm(title, duration, genre, year string, dataFilm *[]map[string]string) {
	film := make(map[string]string)
	film["title"] = title
	film["duration"] = duration
	film["genre"] = genre
	film["year"] = year

	*dataFilm = append(*dataFilm, film)
}
func main() {
	//soal	1
	var jarijari float64 = 14
	var luasLigkaran float64
	var kelilingLingkaran float64

	luaskelilingLingkaran(&jarijari, &luasLigkaran, &kelilingLingkaran)

	fmt.Printf("Luas Lingkaran 		: %f cm\n", luasLigkaran)
	fmt.Printf("Keliling Lingkaran 	: %f cm\n", kelilingLingkaran)
	//soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
	//soal 3
	var buah []string
	tambahDataBuah(&buah, "Jeruk")
	tambahDataBuah(&buah, "Semangka")
	tambahDataBuah(&buah, "Mangga")
	tambahDataBuah(&buah, "Strawberry")
	tambahDataBuah(&buah, "Durian")
	tambahDataBuah(&buah, "Manggis")
	tambahDataBuah(&buah, "Alpukat")

	for i, dataBuah := range buah {
		fmt.Printf("%d. %s \n", i+1, dataBuah)
	}
	//soal 4
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, dataFilms := range dataFilm {
		var no int = i + 1
		fmt.Printf("%d. title: %s\n", no, dataFilms["title"])
		fmt.Printf("   duration: %s\n", dataFilms["duration"])
		fmt.Printf("   genre: %s\n", dataFilms["genre"])
		fmt.Printf("   year: %s\n\n", dataFilms["year"])
	}
}
