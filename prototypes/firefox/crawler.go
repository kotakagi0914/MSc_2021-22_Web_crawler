package main

import (
	"fmt"
	"log"

	"github.com/tebeka/selenium"
)

const (
	targetURL    = "https://recaptcha-demo.appspot.com/recaptcha-v3-request-scores.php"
	seleniumPort = 4444
)

func main() {
	log.Println("Start")

	selenium.SetDebug(true)
	cap := selenium.Capabilities{"browserName": "firefox"}

	// Connect to the WebDriver instance running in a docker container.
	wd, err := selenium.NewRemote(cap, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", seleniumPort))
	if err != nil {
		log.Fatalln("Failed to connect to webDriver: ", err)
	}
	defer wd.Quit()

	// Navigate to the target URL.
	if err := wd.Get(targetURL); err != nil {
		log.Println("Failed to connect to target URL: ", err)
		return
	}

	// Get an HTML tag for the reCAPTCHA step description.
	elem, err := wd.FindElement(selenium.ByCSSSelector, ".step1")
	if err != nil {
		log.Println("Failed to find HTML element: ", err)
		return
	}

	elemText, err := elem.Text()
	if err != nil {
		log.Println("Failed to get tag text: ", err)
		return
	}
	log.Println("Elem text: ", elemText)

	log.Println("End")
}

/*
# Reference
- https://pkg.go.dev/github.com/tebeka/selenium#pkg-overview

# Line Count
- Total:      14
- Reused:     0
- Written:    14
- Referenced: 0
*/
