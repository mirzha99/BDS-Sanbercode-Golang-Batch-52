package soal

import (
	"fmt"
	"sort"
	"time"
)

func tambahDataPhone(phones *[]string) {
	phonelist := []string{"Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo"}
	*phones = append(*phones, phonelist...)
}

func Soal4() {
	fmt.Println("\nSoal 4")
	var phones = []string{}
	tambahDataPhone(&phones)
	sort.Strings(phones)
	for i, phone := range phones {
		fmt.Printf("%v. %v\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}
}
