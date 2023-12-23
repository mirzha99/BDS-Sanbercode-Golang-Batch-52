package helper

import (
	"fmt"
)

// struct soal 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

// interface soal 1 hitungBangunDatar
type HitungBangunDatar interface {
	luas() int
	keliling() int
}

//method bagunan datar
/*	segitiga sama sisi*/
func (s *SegitigaSamaSisi) luas() int {
	return (s.Alas * s.Tinggi) / 2
}
func (s *SegitigaSamaSisi) keliling() int {
	return 3 * s.Alas
}
func HasilHitungSegitigaSamaSisi(hitungBangunDatar HitungBangunDatar) {
	fmt.Println("Segitiga Sama Sisi")
	fmt.Println("Luas Segitiga Sama Sisi :", hitungBangunDatar.luas())
	fmt.Println("Keliling Segitiga Sama Sisi :", hitungBangunDatar.keliling())
	fmt.Println()
}

/*	persegi Panjang*/
func (p *PersegiPanjang) luas() int {
	return p.Panjang * p.Lebar
}
func (p *PersegiPanjang) keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}
func HasilHitungPersegiPanjang(hitungBangunDatar HitungBangunDatar) {
	fmt.Println("Persegi Panjang")
	fmt.Println("Luas Persegi Panjang :", hitungBangunDatar.luas())
	fmt.Println("Keliling Persegi Panjang :", hitungBangunDatar.keliling())
	fmt.Println()
}

// interface soal 1 hitungBangunRuang
type HitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

// Tabung
func (t *Tabung) volume() float64 {
	return 3.14 * t.JariJari * t.JariJari * t.Tinggi
}
func (t *Tabung) luasPermukaan() float64 {
	return 2 * 3.14 * t.JariJari * (t.JariJari + t.Tinggi)
}
func HasilHitungTabung(hitungBangunRuang HitungBangunRuang) {
	fmt.Println("Tabung")
	fmt.Println("Volume Tabung :", hitungBangunRuang.volume())
	fmt.Println("Luas Permukaan Tabung :", hitungBangunRuang.volume())
	fmt.Println()
}

// Balok
func (b *Balok) volume() float64 {
	return float64(b.Panjang) * float64(b.Lebar) * float64(b.Tinggi)
}
func (b *Balok) luasPermukaan() float64 {
	return 2 * (float64(b.Panjang)*float64(b.Lebar) + float64(b.Panjang)*float64(b.Tinggi) + float64(b.Lebar)*float64(b.Tinggi))
}
func HasilHitungBalok(hitungBangunRuang HitungBangunRuang) {
	fmt.Println("Balok")
	fmt.Println("Volume Balok :", hitungBangunRuang.volume())
	fmt.Println("Luas Permukaan Balok :", hitungBangunRuang.luasPermukaan())
	fmt.Println()
}

// struct soal 2
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

// interface soal 2
type phone interface {
	infoPhone() string
}

// struct method soal 2
func (p *Phone) infoPhone() string {
	var color string
	for i, listColor := range p.Colors {
		if len(p.Colors)-1 == i {
			color += listColor
		} else {
			color += listColor + ", "
		}

	}
	return fmt.Sprintf("name : %v\nbrand : %v\nyaer : %v\ncolor : %v", p.Name, p.Brand, p.Year, color)
}

// func soal 2
func GetDataPhone(phone Phone) {
	fmt.Println(phone.infoPhone())
	fmt.Println()
}

// func soal 3
func LuasPersegi(sisi int, kondisi bool) interface{} {
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
