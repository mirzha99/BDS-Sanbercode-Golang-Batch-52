package main

import (
	"fmt"
)

func main() {
	//soal 1
	/*
		Perulangan  i = 1 <= 20 maka akan melakukan perulangan 20 x
	*/
	for i := 1; i <= 20; i++ {
		//Kondisi ini jika i habis di bagi dengan 2 ==0 maka akan menghasilkan nilai genap
		if i%2 == 0 {
			fmt.Printf("%d - %s \n", i, "Santai")
			//Kondisi ini jika i habis di bagi dengan 3 == 0 maka akan menghasilkan nilai kelipatan 3 dengan nilai ganjil
		} else if i%3 == 0 {
			fmt.Printf("%d - %s \n", i, "I Love Coding")
			/*
				Jika kondisi di atas tidak terpenuhi maka menjadi else
				dan kebetulan kondisi di atas adalah genap dan ganjil kelipatan 3
				maka nilai else akan menghasilkan ganjil
			*/
		} else {
			fmt.Printf("%d - %s \n", i, "Berkualitas")
		}
	}
	//soal 2
	/*
		- Perulangan i = 1 <= 7 i++ menngrontrol baris-baris sebanyak 7 baris
		- Perulangan j = i <= j j ++ mengrontrol jumlah karakter ("#") dari nilai j=1 hingga nilai `i`
		- fmt.Print("#") mencetak karakter #
		- fmt.Printf("\n") membuat enter pada setiap perulangan yang pertama
	*/
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Printf("\n")
	}
	//soal 3
	// membuat variabel kalimat dalam bentuk array
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	/**
		Membuat variabel kalimatSplit dengan type string di gunakan dalam perulangan
	untuk mengabungkan string dari perulangan array kalimat.

	Jika tidak mau menggunakan perulangan bisa juga di menggunakan dari pakcake strings.Join(kalimat[:]." ")

	Contoh :
		kalimatSplit := strings.Join(kalimat[:], " ")
	*/

	var kalimatSplit string

	for _, kata := range kalimat {
		kalimatSplit += fmt.Sprintf(`%s `, kata)
	}

	//cara ke2 tanpa perulangan kalimatSplit := strings.Join(kalimat[:], " ")

	slice := []string{kalimatSplit}
	fmt.Println("Slice :", slice)
	fmt.Printf("Printf : [%s]\n", kalimatSplit)

	//soal 4
	/** membuat slice sayuran */
	var sayuran = []string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}
	/**
	Melakukan perulangan dari slice sayuran
		- i adalah index yang di mulai dari angka 0
		- sayur adalah variabel baru yg di as kan dari range sayuran
		  (range bisa di gunakan oleh array, slice ,map)
	*/
	for i, sayur := range sayuran {
		/**
		Disini saya menggunakan fmt.Printf("%d. %s\n", i+1, sayur)
			-%d = mendelkrasikan nilai int
			-%s = mendelkrasikan nilai string
			-\n = membuat enter
			- i+1 = dalam penggunaan index (variabel i) for maka nilai di mulai dari nilai 0
			  untuk memulai dari angka 1 saya menggunakan operator matematika untuk menambah 1 angka
			  sehingga menghasilakan 1. Bayam dan seterusnya jika saya tidak melakuakan  i+1 maka akan
			  menghasilkan 0. Bayam dan seterusnya
		*/
		fmt.Printf("%d. %s\n", i+1, sayur)
	}
	//soal 5
	//membuat variabel satuan dalam bentuk map
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	/*
			var volumeBalok int adalah variabel menampung hasil Volume Balok dan nilai tipe data int adalah 0
		di karenakan dalam rumus perhitungan Volume Balok adalah p*l*t ada operator x (*) jika saya tidak mengubah
		nilai dari variabel volumeBalok menjadi 1 maka nilai default 0 maka hasil operatornya menjadi 0
		karena angka apapun di kali 0 tetap hasilnya 0
	*/
	var volumeBalok int = 1
	//melakukan perulangan dari map satuan
	for value, data := range satuan {
		/*
			-value = adalah key dari map satuan
			-data  = adalah data dari key map satuan
		*/
		fmt.Printf("%s = %d\n", value, data)
		/*
			- volumeBalok *= data ialah Augmented Assignments Operator Matematika
			- hasil dari volumeBalok yang di sebabkan oleh perulangan maka secara otomatis
			  akan di kalikan data yang ada dalam perulangan data map satuan volumeBalok (1 * 7 * 4 * 6 = 168)
		*/
		volumeBalok *= data
	}
	fmt.Println("Volume Balok =", volumeBalok)
}
