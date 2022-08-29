package main

import (
	"log"

	"github.com/sheva0914/MSc_2021-22_Web_crawler/pkg/crawler"
)

const (
	// targetURL = "http://mock-server:8000"
	targetURL    = "http://recaptcha-v3-test.com:8000"
	browserName  = "chrome"
	seleniumPort = 4445
)

func main() {
	log.Println("Start")

	for i := 0; i < 1; i++ {
		if err := crawler.Run(browserName, targetURL, seleniumPort, true, true); err != nil {
			log.Printf("Failed to run crawler on %s: %s\n", browserName, err.Error())
		}
	}

	log.Println("End")
}

/*
# Line Count
- Total:      26
- Reused:     0
- Written:    26
- Referenced: 0
*/
