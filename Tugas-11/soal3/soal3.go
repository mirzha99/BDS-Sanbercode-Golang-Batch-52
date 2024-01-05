package soal3

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func hitungLuasLingkaran(jariJari float64, hasilChan chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	luas := math.Pi * math.Pow(jariJari, 2)
	hasilChan <- luas
}

func hitungKelilingLingkaran(jariJari float64, hasilChan chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	keliling := 2 * math.Pi * jariJari
	hasilChan <- keliling
}

func hitungVolumeTabung(jariJari, tinggi float64, hasilChan chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	volume := math.Pi * math.Pow(jariJari, 2) * tinggi
	hasilChan <- volume
}
func Soal3() {
	jariJariLingkaran := []float64{8, 14, 20}
	tinggiTabung := 10

	var wg sync.WaitGroup

	for _, jariJari := range jariJariLingkaran {
		wg.Add(3)
		hasilLuasChan := make(chan float64)
		hasilKelilingChan := make(chan float64)
		hasilVolumeChan := make(chan float64)

		go hitungLuasLingkaran(jariJari, hasilLuasChan, &wg)
		go hitungKelilingLingkaran(jariJari, hasilKelilingChan, &wg)
		go hitungVolumeTabung(jariJari, float64(tinggiTabung), hasilVolumeChan, &wg)

		go func(jariJari float64) {
			wg.Wait()
			close(hasilLuasChan)
			close(hasilKelilingChan)
			close(hasilVolumeChan)
		}(jariJari)

		go func(jariJari float64) {
			luas := <-hasilLuasChan
			keliling := <-hasilKelilingChan
			volume := <-hasilVolumeChan

			fmt.Printf("Lingkaran dengan jari-jari %.2f:\n", jariJari)
			fmt.Printf("Luas: %.2f\n", luas)
			fmt.Printf("Keliling: %.2f\n", keliling)
			fmt.Printf("Volume tabung: %.2f\n\n", volume)
		}(jariJari)
	}

	// sleep program selama 2 detik untuk menjalankan goroutines
	time.Sleep(2 * time.Second)
}
