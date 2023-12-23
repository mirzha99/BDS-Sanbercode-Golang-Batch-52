package soal

import (
	"fmt"
	"strconv"
)

// func soal 1
func Kalimat(Kalimat string) {
	fmt.Println(Kalimat)
}
func Tahun(Tahun int) {
	t := strconv.Itoa(Tahun)
	fmt.Println("Tahun " + t)
}
func Soal1() {
	fmt.Println("Soal 1")
	defer Tahun(2023)
	Kalimat("Golang Backend Development")
}
