/*
*

maaf anda belum menginput sisi dari segitiga sama sisi
*/
package soal

import (
	"errors"
	"fmt"
	"strconv"
)

func kelilingSegitigaSamaSisi(val int, kondisi bool) (string, error) {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("Message Panic :", msg)
		}
	}()
	keliling := val * val
	if val != 0 {
		if kondisi {
			return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", val, keliling), nil
		} else {
			return strconv.Itoa(keliling), nil
		}
	} else {
		if kondisi {
			return "", errors.New("maaf anda belum menginput sisi dari segitiga sama sisi")
		} else {
			panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
	}
}

func Soal2() {
	dataSisi := []map[string]interface{}{
		{"sisi": 4, "kondisi": true},
		{"sisi": 8, "kondisi": false},
		{"sisi": 0, "kondisi": true},
		{"sisi": 0, "kondisi": false},
	}
	for _, data := range dataSisi {
		result, err := kelilingSegitigaSamaSisi(data["sisi"].(int), data["kondisi"].(bool))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}

}
