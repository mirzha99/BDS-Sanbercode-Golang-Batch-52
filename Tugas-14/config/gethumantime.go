package config

import "time"

func GetHumanTime(epochTime int64) string {
	// Konversi waktu Unix epoch ke time.Time
	timeValue := time.Unix(epochTime, 0)

	// Format waktu menjadi string sesuai keinginan
	humanTime := timeValue.Format("2 January 2006 15:04:05 MST")

	return humanTime
}
