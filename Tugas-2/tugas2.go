package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//soal 1
	kata1 := "Bootcamp"
	kata2 := "Digital"
	kata3 := "Skill"
	kata4 := "Sanbercode"
	kata5 := "Golang"
	/**
		Package fmt.Sprintf untuk mempermudah dalam mengabungkan kalimat dengan karakter %s yang merupakan
	penampungan value string yang di ikuti dengan , variabel
	*/
	kalimatResult := fmt.Sprintf("%s %s %s %s %s", kata1, kata2, kata3, kata4, kata5)
	//Package fmt.Print() untuk mencetak
	fmt.Println(kalimatResult)
	//soal 2
	halo := "Halo Dunia"
	/**
	strings.Replace(halo, "Dunia", "Golang", 1) = mengubah kalimat "Dunia" menjadi "Golang" yang di replace 1 kata
	jika ada kata Dunia lebih 1 maka kata Dunia Pertama menjadi Golang
	*/
	fmt.Println(strings.Replace(halo, "Dunia", "Golang", 1))
	//soal 3
	kataPertama := "saya"
	kataKedua := "senang"
	kataKetiga := "belajar"
	kataKeempat := "golang"

	/**
	Mengganti huruf pertama kataKedua dengan huruf besar
		Variabel upperFrist yang bertujuan untuk mengubah huruf pertama menajadi huruf Besar
	strings.ToUpper(string(kataKedua[0])) = untuk mengubah huruf pertama (indexs 0) menjadi huruf besar
	kataKedua[1:len(kataKedua)+0] = mengambil bagian dari kataKedua mulai dari indeks 1 hingga panjang
	*/
	upperFrist := strings.ToUpper(string(kataKedua[0])) + kataKedua[1:len(kataKedua)+0]
	/**Mengganti huruf terakhir kataKetiga dengan huruf besar
		Variabel upperLast yang bertujuan untuk mengubah huruf terakhir menajadi huruf Besar
	kataKetiga[0:len(kataKetiga)-1] = mengambil bagian dari kataKetiga mulai dari indeks 0 hingga panjang
	di kurangi 1 maka otomatis huruf terakhir tidak akan di tampil.
	strings.ToUpper(string(kataKetiga[len(kataKetiga)-1])) = mengubah huruf besar pada index terakhir
	*/
	upperLast := kataKetiga[0:len(kataKetiga)-1] + strings.ToUpper(string(kataKetiga[len(kataKetiga)-1]))
	// Menggabungkan variabel-variabel tersebut
	hasil := fmt.Sprintf("%s %s %s %s", kataPertama, upperFrist, upperLast, strings.ToUpper(kataKeempat))

	// Menampilkan hasil
	fmt.Println(hasil)
	//soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"
	/**
	Mengubah type variabel string menjadi int
		Package strconv.Atoi(string) mengembalikan 2 return yaitu value dan error
	di bawah ini ada variabel a1 dan _  (underscore) adalah return value error fungsi variabel(underscore)
	fungsi error package strconv.Atoi bisa di gunakan untuk pengecekan apakah nilai value string adalah int ?
	jika ada karakter lain selain angka maka akan terjadi invalid syntax (error) pada kasus ini errornya
	di abaikan mengunakan variabel(_) maka tidak terjadi apa apa dan valuenya tidak akan di return value int
	*/
	a1, _ := strconv.Atoi(angkaPertama)
	a2, _ := strconv.Atoi(angkaKedua)
	a3, _ := strconv.Atoi(angkaKetiga)
	a4, _ := strconv.Atoi(angkaKeempat)
	sumAll := a1 + a2 + a3 + a4
	fmt.Println(sumAll)
	//soal 5
	kalimat := "halo halo bandung"
	angka := 2023
	/**
	strings.ReplaceAll(kalimat, "halo", "hi") = mengubah semua kalimat "halo" menjadi "hi"
	*/
	fmt.Printf("\"%s\" - %d", strings.ReplaceAll(kalimat, "halo", "hi"), angka)
}
