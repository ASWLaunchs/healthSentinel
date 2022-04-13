package models

import "time"

func NowDate() string {
	var timeStr string = time.Now().Format("2006-01-02")
	return timeStr
}
