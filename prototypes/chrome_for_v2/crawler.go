package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

const (
	targetURL    = "https://recaptcha-demo.appspot.com/recaptcha-v2-invisible.php"
	seleniumPort = 4444
)

func main() {
	log.Println("Start")

	selenium.SetDebug(true)
	cap := selenium.Capabilities{"browserName": "chrome"}

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

	// Get an HTML tag for submitting a form and click it.
	submitButtonElem, err := wd.FindElement(selenium.ByCSSSelector, ".g-recaptcha")
	if err != nil {
		log.Println("Failed to find submit button: ", err)
		return
	}
	if err := submitButtonElem.Click(); err != nil {
		log.Println("Failed to click submit button: ", err)
		return
	}

	// Wait for the requests completed.
	time.Sleep(time.Second * 2)

	// Get page source.
	pageSource, err := wd.PageSource()
	if err != nil {
		log.Println("Failed to obtain page source: ", err)
	}
	log.Println("Page source: ", pageSource)

	log.Println("End")
}

/*
# Reference
- https://pkg.go.dev/github.com/tebeka/selenium#pkg-overview

# Line Count
- Total:      57
- Reused:     0
- Written:    43
- Referenced: 14
*/
