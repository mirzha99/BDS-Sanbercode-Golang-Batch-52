package soal

import (
	"flag"
	"fmt"
)

func Soal6() {
	fmt.Println("\nSoal 6")
	panjang := flag.Int64("panjang", 5, "panjang persegi panjang")
	lebar := flag.Int64("lebar", 6, "lebar persegi panjang")

	flag.Parse()
	luas := *panjang * *lebar
	keliling := 2 * (*panjang + *lebar)
	fmt.Printf("Luas persegi panjang dari panjang %v cm dan lebar %v cm adalah %v cm\n", *panjang, *lebar, luas)
	fmt.Printf("Keliling persegi panjang dari panjang %v cm dan lebar %v cm adalah %v cm\n", *panjang, *lebar, keliling)
	// go run . -panjang 6 -lebar 9
}
