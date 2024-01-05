package main

import (
	"fmt"
	"net/http"
	"tugas-12/soal"
)

func main() {
	soal.Soal1()
	soal.Soal2()
	soal.Soal3()
	http.HandleFunc("/soal4", soal.Soal4)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("http://localhost:8080/soal4")
}
