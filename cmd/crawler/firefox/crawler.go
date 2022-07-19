package main

import (
	"fmt"
	"log"

	"github.com/tebeka/selenium"
)

const (
	targetURL     = "http://mock-server:8000"
	seleniumPort  = 4444
	loginUserName = "admin"
	loginPassword = "password"
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

	// Get an input HTML tag for username.
	usernameElem, err := wd.FindElement(selenium.ByCSSSelector, "input[name=\"username\"]")
	if err != nil {
		log.Println("Failed to find username input element: ", err)
		return
	}

	// Input username to the input tag.
	if err := usernameElem.SendKeys(loginUserName); err != nil {
		log.Println("Failed to input login username: ", err)
		return
	}

	// // Make sure that the HTML text is NOT displayed before triggering reCAPTCHA requests.
	// reCAPTCHAScoreElemDisplayed, err := reCAPTCHAScoreElem.IsDisplayed()
	// if err != nil {
	// 	log.Println("Failed to get the flag for the tag being displayed: ", err)
	// 	return
	// }
	// log.Println("reCAPTCHA score element displayed?: ", reCAPTCHAScoreElemDisplayed)

	// // Identify the tag to trigger reCAPTCHA requests and click it.
	// triggerElem, err := wd.FindElement(selenium.ByCSSSelector, ".go")
	// if err != nil {
	// 	log.Println("Failed to find trigger element: ", err)
	// 	return
	// }
	// if err := triggerElem.Click(); err != nil {
	// 	log.Println("Failed to click trigger element: ", err)
	// 	return
	// }

	// // Wait for the requests completed.
	// time.Sleep(time.Second * 2)

	// // Make sure the score tag is displayed.
	// reCAPTCHAScoreElemDisplayed, err = reCAPTCHAScoreElem.IsDisplayed()
	// if err != nil {
	// 	log.Println("Failed to get the flag for the tag being displayed: ", err)
	// 	return
	// }
	// log.Println("reCAPTCHA score element displayed?: ", reCAPTCHAScoreElemDisplayed)

	// // Obtain reCAPTCHA score text.
	// reCAPTCHAScoreText, err := reCAPTCHAScoreElem.Text()
	// if err != nil {
	// 	log.Println("Failed to get the flag for the tag being displayed: ", err)
	// 	return
	// }
	// log.Println("reCAPTCHA score: ", reCAPTCHAScoreText)

	log.Println("End")
}

/*
# Reference
- https://pkg.go.dev/github.com/tebeka/selenium#pkg-overview

# Line Count
- Total:      81
- Reused:     0
- Written:    67
- Referenced: 14
*/
