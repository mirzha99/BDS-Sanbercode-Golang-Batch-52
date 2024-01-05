package soal4

import (
	"fmt"
	"sync"
)

func hitungLuasPersegiPanjang(panjang, lebar float64, hasilChan chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	luas := panjang * lebar
	hasilChan <- luas
}

func hitungKelilingPersegiPanjang(panjang, lebar float64, hasilChan chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	keliling := 2 * (panjang + lebar)
	hasilChan <- keliling
}

func hitungVolumeBalok(panjang, lebar, tinggi float64, hasilChan chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	volume := panjang * lebar * tinggi
	hasilChan <- volume
}

func Soal4() {
	panjangPersegi := 4.0
	lebarPersegi := 8.0
	tinggiBalok := 5.0

	var wg sync.WaitGroup

	wg.Add(3)
	hasilLuasChan := make(chan float64)
	hasilKelilingChan := make(chan float64)
	hasilVolumeChan := make(chan float64)

	go hitungLuasPersegiPanjang(panjangPersegi, lebarPersegi, hasilLuasChan, &wg)
	go hitungKelilingPersegiPanjang(panjangPersegi, lebarPersegi, hasilKelilingChan, &wg)
	go hitungVolumeBalok(panjangPersegi, lebarPersegi, tinggiBalok, hasilVolumeChan, &wg)

	go func() {
		wg.Wait()
		close(hasilLuasChan)
		close(hasilKelilingChan)
		close(hasilVolumeChan)
	}()

	for i := 0; i < 3; i++ {
		select {
		case luas := <-hasilLuasChan:
			fmt.Printf("Luas Persegi Panjang: %.2f\n", luas)
		case keliling := <-hasilKelilingChan:
			fmt.Printf("Keliling Persegi Panjang: %.2f\n", keliling)
		case volume := <-hasilVolumeChan:
			fmt.Printf("Volume Balok: %.2f\n", volume)
		}
	}
}
