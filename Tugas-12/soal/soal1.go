package soal

import (
	"encoding/json"
	"fmt"
	"tugas-12/data"
)

func Soal1() {
	var dataDevice []map[string]string
	err := json.Unmarshal([]byte(data.DeviceData()), &dataDevice)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	var dataFilter []map[string]string
	var total int // deafult value 0
	for _, dd := range dataDevice {
		if dd["vendor"] == "Samsung" {
			dataFilter = append(dataFilter, dd)
			total = len(dataFilter)
		}
	}
	fmt.Println(dataFilter)
	fmt.Println("Total : ", total)
}
