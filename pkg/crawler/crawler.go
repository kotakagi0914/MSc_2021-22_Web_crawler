package crawler

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

const (
	seleniumPort  = 4444
	loginUserName = "admin"
	loginPassword = "password"
)

func Run(browserName, targetURL string) error {
	selenium.SetDebug(true)
	cap := selenium.Capabilities{"browserName": browserName}

	// Connect to the WebDriver instance running in a docker container.
	wd, err := selenium.NewRemote(cap, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", seleniumPort))
	if err != nil {
		return fmt.Errorf("[crawler.Run()] Failed to connect to webDriver: %v", err)
	}
	defer wd.Quit()

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

	// Click login button.
	loginButtonElem, err := wd.FindElement(selenium.ByCSSSelector, "input[name=\"submit\"]")
	if err != nil {
		return fmt.Errorf("[crawler.Run()] Failed to find login button element: %v", err)
	}
	if err := loginButtonElem.Click(); err != nil {
		return fmt.Errorf("[crawler.Run()] Failed to click login button element: %v", err)
	}

	// Wait for the requests completed.
	time.Sleep(time.Second * 2)

	// Obtain login result elements.
	loginResulstsElem, err := wd.FindElement(selenium.ByCSSSelector, "div[name=\"login-result\"]")
	if err != nil {
		return fmt.Errorf("[crawler.Run()] Failed to find login result elements: %v", err)
	}

	// Get login results text and show it.
	loginResult, err := loginResulstsElem.Text()
	if err != nil {
		return fmt.Errorf("[crawler.Run()] Failed to get login result text: %v", err)
	}
	log.Println("Login result: ", loginResult)

	return nil
}

/*
# Reference
- https://pkg.go.dev/github.com/tebeka/selenium#pkg-overview

# Line Count
- Total:      93
- Reused:     0
- Written:    80
- Referenced: 13
*/
