package main

import (
	"fmt"
	"strings"
)

//func soal 1
/**
Cara Membuat Function Predefined return value
	func namaFunct(var1,var2,var3 tipe data) (int,error){
		result := var1 + var2 +vvar3
		return result,nil
	}
*/
func luasPersegiPanjang(p, l int) int {
	//mengembalikan nilai dari variabel p di kali dengan variabel l, 12 * 4 = 48
	return p * l
}

func kelilingPersegiPanjang(p, l int) int {
	//mengembalikan hasil nilai dari perhitungan 2x(p+l), 2 × (12+4) = 32 || 12+12+4+4 = 32
	return 2 * (p + l)
}
func volumeBalok(p, l, t int) int {
	//mengembalikan hasil p x l x t, 12 * 4 * 8 = 384
	return p * l * t
}

// func soal 2
func introduce(name, genders, jobs, age string) string {
	/**
	-	variabel gender berfungsi untuk mengembalikan nilai string "Pak" atau "Bu" berdasarakan jenis kelamin
	-	strings.ToLower mengubah huruf besar menjadi haruf kecil jika ada pengguna menginputkan huruf
		besar tidak akan terjadi false
	-	jika jenis kelamin tidak ada di kondisi di atas maka else akan meretrun Tolak LGBT :v
	*/
	var gender string
	if strings.ToLower(genders) == "laki-laki" || strings.ToLower(genders) == "cowok" {
		gender = "Pak"
	} else if strings.ToLower(genders) == "perempuan" || strings.ToLower(genders) == "cewek" {
		gender = "Bu"
	} else {
		return "Tolak LGBT :v"
	}
	result := fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", gender, name, jobs, age)
	return result
}

// func soal 3
func buahFavorit(name string, buah ...string) string {
	var buahs string
	strings.Join(buah, "")
	for i, dataBuah := range buah {
		/**
		* len(buah)-1 mengamabil index terkahir dari slice buah
		* jika i belum sampai ke index terkahir maka variabel buahs akan mereturn "namabuah", (ada koma)
		* jika i sudah sampai ke index terkahir maka variabel buahs akan mereturn "namabuah" (tanpa koma)
		 */
		if i != len(buah)-1 {
			buahs += fmt.Sprintf(`"%s", `, dataBuah)
		} else {
			buahs += fmt.Sprintf(`"%s"`, dataBuah)
		}

	}
	return fmt.Sprintf(`halo nama saya %s dan buah favorit saya adalah %s`, name, buahs)
}
func main() {
	//soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	//soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	//soal 4
	var dataFilm = []map[string]string{}
	// buatlah closure function disini

	/**
	Membuat closure function tambahDataFilm dan anonymous func dengan variadic input ...string
	*/
	tambahDataFilm := func(input ...string) {
		//memvalidasi input, jika input != 4 maka closure function menghasilkan map[msg:Input Tidak Valid]
		if len(input) != 4 {
			dataFilm = append(dataFilm, map[string]string{"msg": "Input Tidak Valid"})
		} else {
			//membuat map baru dengan func make
			filmData := make(map[string]string)
			//menapung keyMap dengan slice untuk membuat key pada map filmData
			keyMap := []string{"title", "jam", "genre", "tahun"}
			//membuat maps baru dan input data dengan menggunakan perulangan for
			for i, keys := range keyMap {
				/**
				Proses add data film dengan Membuat Map Baru
					- i = berfungsi memasukan nilai index sesuai input variadic
					- keys = mendeklarasikan key map pada map filmData
					- filmData["title"] = "LOTR" dan seterusnya
				*/
				filmData[keys] = input[i]
			}
			/**
			Menambah data maps filmData pada maps dataFilm
				jika melakukan perulangan map dataFilm maka secara otomatis data maps filmData akan di tampil
			di karenakan maps dataFilm sudah di append dengan maps filmData
			*/
			dataFilm = append(dataFilm, filmData)
		}

	}
	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")
	tambahDataFilm("Dua Garis Biru", "2 jam setengah", "drama", "2021", "test") // map[msg:Input Tidak Valid]

	for _, item := range dataFilm {
		fmt.Println(item)
	}
	//maaf tergangu dengan komentar :v
}
