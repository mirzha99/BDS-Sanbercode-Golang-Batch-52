package soal

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title, Desc, Author string
	ReleaseYear         int
}

func Soal3() {
	data := []Book{
		{Title: "Belajar Golang Dasar", Desc: "Tutorial Golang", Author: "Agung Noval", ReleaseYear: 2023},
		{Title: "Atomic Habits", Desc: "Mengubah hal kecil menjadi lebih baik", Author: "Jemes Clear", ReleaseYear: 2016},
		{Title: "Managemen Proyek", Desc: "Perancangan Managenet Proyek", Author: "D Rumi", ReleaseYear: 2012},
		{Title: "Berani Tidak di sukai", Desc: "Menjadi Diri Sendiri", Author: "Azrim", ReleaseYear: 2024},
		{Title: "Ada Apa Dengan Cinta", Desc: "Cinta", Author: "Dea", ReleaseYear: 2017},
	}

	json, err := json.Marshal(&data)

	if err != nil {
		fmt.Println("Error encode Json :", err.Error())
	}

	fmt.Println(string(json))
}
