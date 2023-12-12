package main

import (
	"fmt"
	"strconv"
)

func main() {
	//soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	panjangPersegiPanjangs, _ := strconv.Atoi(panjangPersegiPanjang)
	lebarPersegiPanjangs, _ := strconv.Atoi(lebarPersegiPanjang)
	alasSegitigas, _ := strconv.Atoi(alasSegitiga)
	tinggiSegitigas, _ := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int = panjangPersegiPanjangs * lebarPersegiPanjangs
	var kelilingPersegiPanjang int = 2 * (panjangPersegiPanjangs + lebarPersegiPanjangs)
	var luasSegitiga int = (alasSegitigas * tinggiSegitigas) / 2

	fmt.Println(luasPersegiPanjang)
	fmt.Println(kelilingPersegiPanjang)
	fmt.Println(luasSegitiga)

	//soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50
	//if-else
	if nilaiJohn >= 80 {
		fmt.Println("indeksnya A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("indeksnya B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("indeksnya C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("indeksnya D")
	} else if nilaiJohn < 50 {
		fmt.Println("indeksnya E")
	}
	//switch - case
	switch {
	case nilaiDoe >= 80:
		fmt.Println("indeksnya A")
	case nilaiDoe >= 70 && nilaiDoe < 80:
		fmt.Println("indeksnya B")
	case nilaiDoe >= 60 && nilaiDoe < 70:
		fmt.Println("indeksnya C")
	case nilaiDoe >= 50 && nilaiDoe < 60:
		fmt.Println("indeksnya D")
	case nilaiDoe < 50:
		fmt.Println("indeksnya E")
	}
	//soal 3

	var tanggal = 31
	var bulan = 07
	var tahun = 1999
	namabulan := ""
	switch bulan {
	case 1:
		namabulan = "Januari"
	case 2:
		namabulan = "Febuari"
	case 3:
		namabulan = "Maret"
	case 4:
		namabulan = "April"
	case 5:
		namabulan = "Mei"
	case 6:
		namabulan = "Juni"
	case 7:
		namabulan = "Juli"
	case 8:
		namabulan = "Agustus"
	case 9:
		namabulan = "September"
	case 10:
		namabulan = "Oktober"
	case 11:
		namabulan = "November"
	case 12:
		namabulan = "Desember"
	default:
		namabulan = "Tidak Ada Bulan " + strconv.Itoa(bulan)
	}

	fmt.Printf("%d %s %d \n", tanggal, namabulan, tahun)
	//soal 4
	/**
	Baby boomer, kelahiran 1944 s.d 1964
	Generasi X, kelahiran 1965 s.d 1979
	Generasi Y (Millenials), kelahiran 1980 s.d 1994
	Generasi Z, kelahiran 1995 s.d 2015
	*/
	tahunkelahiran := tahun

	if tahunkelahiran >= 1944 && tahunkelahiran < 1964 {
		fmt.Println("Tahun " + strconv.Itoa(tahunkelahiran) + " Baby Boomer")
	} else if tahunkelahiran >= 1965 && tahunkelahiran < 1979 {
		fmt.Println("Tahun " + strconv.Itoa(tahunkelahiran) + " Generasi X")
	} else if tahunkelahiran >= 1980 && tahunkelahiran < 1994 {
		fmt.Println("Tahun " + strconv.Itoa(tahunkelahiran) + " Generasi Y (Millenials)")
	} else if tahunkelahiran >= 1995 && tahunkelahiran < 2015 {
		fmt.Println("Tahun " + strconv.Itoa(tahunkelahiran) + " Generasi Z")
	} else {
		fmt.Println("Generasi Tahun " + strconv.Itoa(tahunkelahiran) + " Tidak Diketahui")
	}
}
