package soal

import (
	"fmt"
	"math"
)

type jariLingkran struct {
	jarijari int
}
type hitung interface {
	getJariJari() int
	luasLingkaran() float64
	kelilingLingkaran() float64
}

func (jl *jariLingkran) getJariJari() int {
	return jl.jarijari
}
func (jl *jariLingkran) luasLingkaran() float64 {
	luas := math.Pi * math.Pow(float64(jl.jarijari), 2)
	return math.Round(luas*10) / 10
}
func (jl *jariLingkran) kelilingLingkaran() float64 {
	keliling := 2 * math.Pi * float64(jl.jarijari)
	return math.Round(keliling*10) / 10
}

func hitungLingkaran(hitung hitung) {
	j := hitung.getJariJari()
	fmt.Printf("Luas lingkaran jari-jari %v cm adalah %v cm\n", j, hitung.luasLingkaran())
	fmt.Printf("Keliling lingkaran jari-jari %v cm adalah %v cm\n\n", j, hitung.kelilingLingkaran())
}

func Soal5() {
	fmt.Println("\nSoal 5")
	var jarijari = []jariLingkran{
		{jarijari: 7},
		{jarijari: 10},
		{jarijari: 15},
	}
	for _, datajarijari := range jarijari {
		hitungLingkaran(&datajarijari)
	}
}
