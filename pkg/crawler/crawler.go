package crawler

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

const (
	loginUserName = "admin"
	loginPassword = "password"
)

func Run(browserName, targetURL string, portNum int) error {
	// selenium.SetDebug(true)
	cap := selenium.Capabilities{"browserName": browserName}

	// Connect to the WebDriver instance running in a docker container.
	wd, err := selenium.NewRemote(cap, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", portNum))
	if err != nil {
		return fmt.Errorf("[crawler.Run()] Failed to connect to webDriver: %v", err)
	}
	defer wd.Quit()

	for i := 0; i < 10; i++ {
		// Navigate to the target URL.
		if err := wd.Get(targetURL); err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to connect to target URL: %v", err)
		}

		// Get an input HTML tag for username.
		usernameElem, err := wd.FindElement(selenium.ByCSSSelector, "input[name=\"username\"]")
		if err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to find username input element: %v", err)
		}

		// Input username to the input tag.
		if err := usernameElem.SendKeys(loginUserName); err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to input login username: %v", err)
		}

		// Get an input HTML tag for password.
		passwordElem, err := wd.FindElement(selenium.ByCSSSelector, "input[name=\"password\"]")
		if err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to find password input element: %v", err)
		}

		// Input password to the input tag.
		if err := passwordElem.SendKeys(loginPassword); err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to input login password: %v", err)
		}

		// Get an checkbox HTML tag anc click it.
		checkboxElem, err := wd.FindElement(selenium.ByCSSSelector, "input[name=\"checkbox\"]")
		if err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to find checkbox input element: %v", err)
		}
		if err := checkboxElem.Click(); err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to click checkbox element: %v", err)
		}

		// Click login button.
		loginButtonElem, err := wd.FindElement(selenium.ByCSSSelector, "input[name=\"submit\"]")
		if err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to find login button element: %v", err)
		}
		if err := loginButtonElem.Click(); err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to click login button element: %v", err)
		}

		// Wait for the requests completed.
		time.Sleep(time.Second * 1)

		// Obtain login result elements.
		// When failed, sleep and retry up tp 10 times.
		var loginResulstsElem selenium.WebElement
		for i := 0; i < 10; i++ {
			loginResulstsElem, err = wd.FindElement(selenium.ByCSSSelector, "div[name=\"login-result\"]")
			if err == nil {
				break
			}
			log.Println("[crawler.Run()] Sleep and retry to get login result")
			time.Sleep(time.Second * 1)
		}
		if err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to find login result elements: %v", err)
		}

		// Get login results text and show it.
		loginResult, err := loginResulstsElem.Text()
		if err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to get login result text: %v", err)
		}
		log.Println("Login result: ", loginResult)
	}

	return nil
}

/*
# Reference
- https://pkg.go.dev/github.com/tebeka/selenium#pkg-overview

# Line Count
- Total:      100
- Reused:     0
- Written:    66
- Referenced: 34
*/
