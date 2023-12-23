package soal

import "fmt"

func tambahAngka(value int, angka *int) {
	*angka += value
}

func cetakAngka(total *int) {
	fmt.Println("Total Angka :", *total)
}
func Soal3() {
	fmt.Println("\nSoal 3")
	angka := 1

	defer cetakAngka(&angka)

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)
}
