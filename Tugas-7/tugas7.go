package main

import "fmt"

//struct soal 1
type buah struct {
	nama, warna string
	adaBijinya  bool
	harga       int64
}

//struct soal 2
type segitiga struct {
	alas, tinggi int
}
type persegi struct {
	sisi int
}
type persegiPanjang struct {
	panjang, lebar int
}

//struct soal 3
type phone struct {
	name, brand string
	year        int
	color       []string
}

//struct soal 4
type movie struct {
	title, genre   string
	duration, year int
}

//method soal 2
//membuat struct method
func (s *segitiga) luasSegitiga() float64 {
	return 0.5 * float64(s.alas) * float64(s.tinggi)
}
func (p *persegi) luasPersegi() int {
	return p.sisi * p.sisi
}
func (pp *persegiPanjang) luasPersegiPanjang() int {
	return pp.panjang * pp.lebar
}

//method soal 3
func (p *phone) tambahWarna(warnaBaru string) {
	p.color = append(p.color, warnaBaru)
}

//method soal 4
func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	movie := movie{
		title:    title,
		duration: duration,
		genre:    genre,
		year:     year,
	}

	movie.duration = duration / 60
	*dataFilm = append(*dataFilm, movie)
}
func main() {
	//soal 1
	//membuat slice dengan struct buah
	buah := []buah{
		{nama: "Nanas", warna: "Kuning", adaBijinya: false, harga: 9000},
		{nama: "Jeruk", warna: "Oranye", adaBijinya: true, harga: 8000},
		{nama: "Semangka", warna: "Hijau & Merah", adaBijinya: true, harga: 10000},
		{nama: "Pisang", warna: "Kuning", adaBijinya: false, harga: 9000},
	}
	//melakukan perulangan dari slice buah struct
	for _, dataBuah := range buah {
		fmt.Println(dataBuah)
	}
	//soal 2
	segitiga := segitiga{alas: 5, tinggi: 8}
	fmt.Printf("\nLuas segitiga dari alas %v dan tinggi %v adalah %v \n", segitiga.alas, segitiga.tinggi, segitiga.luasSegitiga())
	persegi := persegi{sisi: 9}
	fmt.Printf("Luas Persegi dari sisi %v adalah %v\n", persegi.sisi, persegi.luasPersegi())
	persegipanjang := persegiPanjang{panjang: 5, lebar: 7}
	fmt.Printf("Luas Persegi Panjang dari panjang %v dan lebar %v adalah %v \n\n", persegipanjang.panjang, persegipanjang.lebar, persegipanjang.luasPersegiPanjang())
	//soal 3
	PocoF1 := phone{
		name:  "Pocophone F1",
		brand: "Xiaomi",
		year:  2018,
		color: []string{
			"Black",
			"Gray",
		},
	}
	PocoF1.tambahWarna("Red")
	PocoF1.tambahWarna("Yellow")
	fmt.Println(PocoF1)
	fmt.Println()
	//soal 4
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, dataFilms := range dataFilm {
		var no int = i + 1
		fmt.Printf("%d. title: %s\n", no, dataFilms.title)
		fmt.Printf("   duration: %d jam\n", dataFilms.duration)
		fmt.Printf("   genre: %s\n", dataFilms.genre)
		fmt.Printf("   year: %d\n\n", dataFilms.year)
	}
}
