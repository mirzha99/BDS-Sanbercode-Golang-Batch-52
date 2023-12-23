package main

import (
	"fmt"
	"strconv"
)

// struct soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

// interface soal 1 hitungBangunDatar
type hitungBangunDatar interface {
	luas() int
	keliling() int
}

//method bagunan datar
/*	segitiga sama sisi*/
func (s *segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}
func (s *segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}
func hasilHitungSegitigaSamaSisi(hitungBangunDatar hitungBangunDatar) {
	fmt.Println("Segitiga Sama Sisi")
	fmt.Println("Luas Segitiga Sama Sisi :", hitungBangunDatar.luas())
	fmt.Println("Keliling Segitiga Sama Sisi :", hitungBangunDatar.keliling())
	fmt.Println()
}

/*	persegi panjang*/
func (p *persegiPanjang) luas() int {
	return p.panjang * p.lebar
}
func (p *persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}
func hasilHitungPersegiPanjang(hitungBangunDatar hitungBangunDatar) {
	fmt.Println("Persegi Panjang")
	fmt.Println("Luas Persegi Panjang :", hitungBangunDatar.luas())
	fmt.Println("Keliling Persegi Panjang :", hitungBangunDatar.keliling())
	fmt.Println()
}

// interface soal 1 hitungBangunRuang
type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

// tabung
func (t *tabung) volume() float64 {
	return 3.14 * t.jariJari * t.jariJari * t.tinggi
}
func (t *tabung) luasPermukaan() float64 {
	return 2 * 3.14 * t.jariJari * (t.jariJari + t.tinggi)
}
func hasilHitungTabung(hitungBangunRuang hitungBangunRuang) {
	fmt.Println("Tabung")
	fmt.Println("Volume Tabung :", hitungBangunRuang.volume())
	fmt.Println("Luas Permukaan Tabung :", hitungBangunRuang.volume())
	fmt.Println()
}

// balok
func (b *balok) volume() float64 {
	return float64(b.panjang) * float64(b.lebar) * float64(b.tinggi)
}
func (b *balok) luasPermukaan() float64 {
	return 2 * (float64(b.panjang)*float64(b.lebar) + float64(b.panjang)*float64(b.tinggi) + float64(b.lebar)*float64(b.tinggi))
}
func hasilHitungBalok(hitungBangunRuang hitungBangunRuang) {
	fmt.Println("Balok")
	fmt.Println("Volume Balok :", hitungBangunRuang.volume())
	fmt.Println("Luas Permukaan Balok :", hitungBangunRuang.luasPermukaan())
	fmt.Println()
}

// struct soal 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

// interface soal 2
type Phone interface {
	infoPhone() string
}

// struct method soal 2
func (p *phone) infoPhone() string {
	var color string
	for i, listColor := range p.colors {
		if len(p.colors)-1 == i {
			color += listColor
		} else {
			color += listColor + ", "
		}

	}
	return fmt.Sprintf("name : %v\nbrand : %v\nyaer : %v\ncolor : %v", p.name, p.brand, p.year, color)
}

// func soal 2
func getDataPhone(phone Phone) {
	fmt.Println(phone.infoPhone())
	fmt.Println()
}

// func soal 3
func luasPersegi(sisi int, kondisi bool) interface{} {
	hasil := sisi * sisi

	if sisi != 0 {
		if kondisi {
			return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, hasil)
		} else {
			return hasil
		}
	} else {
		if kondisi {
			return "Maaf anda belum menginput sisi dari persegi"
		} else {
			return nil
		}
	}
}
func main() {
	//soal 1
	//bagunan datar
	fmt.Println("Hitung Bagun Datar")
	var segitigaSamaSisi hitungBangunDatar = &segitigaSamaSisi{alas: 8, tinggi: 9}
	hasilHitungSegitigaSamaSisi(segitigaSamaSisi)
	var persegiPanjang hitungBangunDatar = &persegiPanjang{panjang: 3, lebar: 9}
	hasilHitungPersegiPanjang(persegiPanjang)
	//bangun ruang
	fmt.Println("Hitung Bagun Ruang")
	var tabung hitungBangunRuang = &tabung{jariJari: 7, tinggi: 6}
	hasilHitungTabung(tabung)
	var balok hitungBangunRuang = &balok{panjang: 12, lebar: 15, tinggi: 25}
	hasilHitungBalok(balok)

	//soal 2
	samsungPhone := phone{name: "Samsung Galaxy 20", brand: "Samsung", year: 2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"}}
	getDataPhone(&samsungPhone)

	//soal 3
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))
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
