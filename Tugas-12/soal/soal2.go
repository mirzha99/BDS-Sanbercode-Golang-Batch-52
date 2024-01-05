package soal

import (
	"encoding/json"
	"fmt"
	"tugas-12/data"
)

type DeviceData struct {
	Model  string
	Vendor string
}

func Soal2() {
	var dataDevice []DeviceData
	err := json.Unmarshal([]byte(data.DeviceData()), &dataDevice)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	var dataDeviceFilter []DeviceData
	var total int
	for _, dd := range dataDevice {
		if dd.Vendor == "Sony" {
			dataDeviceFilter = append(dataDeviceFilter, dd)
			total = len(dataDeviceFilter)
		}
	}

	fmt.Println(dataDeviceFilter)
	fmt.Println("Total : ", total)
}
