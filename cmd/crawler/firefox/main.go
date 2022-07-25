package main

import (
	"log"

	"github.com/sheva0914/MSc_2021-22_Web_crawler/pkg/crawler"
)

const (
	targetURL   = "http://mock-server:8000"
	browserName = "firefox"
)

func main() {
	log.Println("Start")

	if err := crawler.Run(browserName, targetURL); err != nil {
		log.Fatalf("Failed to run crawler on %s: %s\n", browserName, err.Error())
	}

	log.Println("End")
}

/*
# Line Count
- Total:      22
- Reused:     0
- Written:    22
- Referenced: 0
*/
