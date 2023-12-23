package main

import (
	"fmt"
	"strconv"
	"tugas-9/helper"
)

func main() {
	//soal 1
	//bagunan datar
	fmt.Println("Hitung Bagun Datar")
	var segitigaSamaSisi helper.HitungBangunDatar = &helper.SegitigaSamaSisi{Alas: 8, Tinggi: 9}
	helper.HasilHitungSegitigaSamaSisi(segitigaSamaSisi)
	var persegiPanjang helper.HitungBangunDatar = &helper.PersegiPanjang{Panjang: 3, Lebar: 9}
	helper.HasilHitungPersegiPanjang(persegiPanjang)
	//bangun ruang
	fmt.Println("Hitung Bagun Ruang")
	var tabung helper.HitungBangunRuang = &helper.Tabung{JariJari: 7, Tinggi: 6}
	helper.HasilHitungTabung(tabung)
	var balok helper.HitungBangunRuang = &helper.Balok{Panjang: 12, Lebar: 15, Tinggi: 25}
	helper.HasilHitungBalok(balok)

	//soal 2
	samsungPhone := helper.Phone{Name: "Samsung Galaxy 20", Brand: "Samsung", Year: 2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"}}
	helper.GetDataPhone(samsungPhone)

	//soal 3
	fmt.Println(helper.LuasPersegi(4, true))
	fmt.Println(helper.LuasPersegi(8, false))
	fmt.Println(helper.LuasPersegi(0, true))
	fmt.Println(helper.LuasPersegi(0, false))
	//soal 4
	var prefix interface{} = "hasil penjumlahan dari"
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	prefix, ok1 := prefix.(string)
	AngkaPertama, ok2 := kumpulanAngkaPertama.([]int)
	AngkaKedua, ok3 := kumpulanAngkaKedua.([]int)
	if !ok1 || !ok2 || !ok3 {
		fmt.Println("Type Assertion Gagal !")
		return
	}
	gabung := append(AngkaPertama, AngkaKedua...)
	var angka string
	var hasil int
	for i, dataAngka := range gabung {
		if len(gabung)-1 != i {
			angka += strconv.Itoa(dataAngka) + "+"
		} else {
			angka += strconv.Itoa(dataAngka) + " = "
		}
		hasil += dataAngka
	}
	fmt.Printf("%s %s%d", prefix, angka, hasil)
}
