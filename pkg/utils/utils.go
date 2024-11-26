package utils

import "time"

func GetTodaysDate() string {
	today := time.Now()
	formattedDate := today.Format("2006-01-02")
	return formattedDate
}
