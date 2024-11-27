package utils

import (
	"os"
	"time"
)

func GetTodaysDate() string {
	today := time.Now()
	formattedDate := today.Format("2006-01-02")
	return formattedDate
}

func CheckDebug() bool {
	debug := os.Getenv("SB_DEBUG")
	return debug == "true"
}
