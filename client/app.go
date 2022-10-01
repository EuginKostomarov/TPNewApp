package TP_client

import "time"

func init() {
	InitilizeApp()
	go StartParseData()
	go InitDriver()
	time.Sleep(300 * time.Millisecond)
}

func StartApp() {
	CreateNewAuthWindow()
}
