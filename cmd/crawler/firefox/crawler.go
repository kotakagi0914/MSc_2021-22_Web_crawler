package main

import (
	"fmt"
	"log"
	"time"

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

	// Get an input HTML tag for password.
	passwordElem, err := wd.FindElement(selenium.ByCSSSelector, "input[name=\"password\"]")
	if err != nil {
		log.Println("Failed to find password input element: ", err)
		return
	}

	// Input password to the input tag.
	if err := passwordElem.SendKeys(loginPassword); err != nil {
		log.Println("Failed to input login password: ", err)
		return
	}

	// Click login button.
	loginButtonElem, err := wd.FindElement(selenium.ByCSSSelector, "input[name=\"submit\"]")
	if err != nil {
		log.Println("Failed to find login button element: ", err)
		return
	}
	if err := loginButtonElem.Click(); err != nil {
		log.Println("Failed to click login button element: ", err)
		return
	}

	// Wait for the requests completed.
	time.Sleep(time.Second * 2)

	// Obtain login result elements.
	loginResulstsElem, err := wd.FindElement(selenium.ByCSSSelector, "div[name=\"login-result\"]")
	if err != nil {
		log.Println("Failed to find login result elements: ", err)
		return
	}

	// Get login results text and show it.
	loginResult, err := loginResulstsElem.Text()
	if err != nil {
		log.Println("Failed to get login result text: ", err)
		return
	}
	log.Println("Login result: ", loginResult)

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
