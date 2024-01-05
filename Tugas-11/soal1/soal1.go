package soal1

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

func mengirimDataPhoneKeChannel(ch chan string, wg *sync.WaitGroup) {
	defer close(ch)
	defer wg.Done()

	for i, phone := range phones {
		//send data phone ke channel
		ch <- strconv.Itoa(i+1) + ". " + phone
		time.Sleep(time.Second)
	}
}
func showDataPhoneDariChannel(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for phone := range ch {
		// Menampilkan data phone dari channel
		fmt.Println(phone)
	}
}
func Soal1() {
	// Mengurutkan data phones
	sort.Strings(phones)
	//Membuat channel
	ch := make(chan string)
	//Menggunakan WaitGroup untuk menunggu  goroutine
	var wg sync.WaitGroup
	//goroutine  mengirim data ke channel
	wg.Add(1)
	go mengirimDataPhoneKeChannel(ch, &wg)
	//Menjalankan goroutine untuk menampilkan data dari channel
	wg.Add(1)
	go showDataPhoneDariChannel(ch, &wg)
	// Menunggu semua goroutine selesai sebelum program berakhir
	wg.Wait()
}
