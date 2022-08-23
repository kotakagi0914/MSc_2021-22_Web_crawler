package crawler

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/sheva0914/selenium"
	"github.com/sheva0914/selenium/chrome"
)

const (
	loginUserName = "admin"
	loginPassword = "password"
	typedString   = "abcdefghij"
	duration      = 250
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))

	userAgents = []string{
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0",                                                                             // Firefox on Mac
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36]",                                           // Chrome on Mac
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.6 Safari/605.1.15]",                                           // Safari on Mac
		"Mozilla/5.0 (X11; Ubuntu; Linux aarch64; rv:102.0) Gecko/20100101 Firefox/102.0",                                                                                  // Firefox on Ubuntu
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246",                                  // Edge on Windows 10
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36",                                                    // Chrome on Windows 7
		"Mozilla/5.0 (Linux; Android 12; Pixel 6 Build/SD1A.210817.023; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/94.0.4606.71 Mobile Safari/537.36",   // Chrome on Google Pixel 6
		"Mozilla/5.0 (Linux; Android 12; SM-S906N Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/80.0.3987.119 Mobile Safari/537.36", // Chrome on Samsung Galaxy S22
		"Mozilla/5.0 (iPhone14,3; U; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Mobile/19A346 Safari/602.1",                   // Safari on iPhone 13 Pro Max
		"Mozilla/5.0 (iPhone12,1; U; CPU iPhone OS 13_0 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Mobile/15E148 Safari/602.1",                   // Safari on iPhone 11
	}
	screenResolutions = []string{
		"1920,1080",
		"1366,768",
		"1280,1024",
		"1280,800",
		"1024,768	",
	}
	languages = []string{
		"en-GB",
		"en-US",
		"en-CA",
		"en-In",
		"ja-JP",
		"de-DE",
		"tl-PH",
		"pt-BR",
		"es-AR",
	}
)

func getRandomUserAgent() string {
	return userAgents[r.Intn(len(userAgents))]
}

func getRandomScreenResolution() string {
	return screenResolutions[r.Intn(len(screenResolutions))]
}

func getRandomLanguage() string {
	return languages[r.Intn(len(languages))]
}

func makeArgsForBrowserOptions() []string {
	return []string{
		fmt.Sprintf("user-agent=%s", getRandomUserAgent()),
		fmt.Sprintf("window-size=%s", getRandomScreenResolution()),
		fmt.Sprintf("accept-lang=%s", getRandomLanguage()),
	}
}

func Run(browserName, targetURL string, portNum int, randomParamsEnabled bool) error {
	// selenium.SetDebug(true)
	cap := selenium.Capabilities{"browserName": browserName}
	// Set random parameters for Chrome browser.
	if randomParamsEnabled {
		args := makeArgsForBrowserOptions()
		log.Println("Args: ", args)
		if browserName == "chrome" {
			cap.AddChrome(chrome.Capabilities{Args: args})
		}
	}

	// Connect to the WebDriver instance running in a docker container.
	wd, err := selenium.NewRemote(cap, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", portNum))
	if err != nil {
		return fmt.Errorf("[crawler.Run()] Failed to connect to webDriver: %v", err)
	}
	defer wd.Quit()

	for i := 0; i < 1; i++ {
		// Navigate to the target URL.
		if err := wd.Get(targetURL); err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to connect to target URL: %v", err)
		}

		// Sleep for a bit when loading a page.
		time.Sleep(time.Second * 1)

		// Get an input HTML element for username.
		usernameElem, err := wd.FindElement(selenium.ByCSSSelector, "input[name=\"username\"]")
		if err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to find username input element: %v", err)
		}

		// Get the location of the username element.
		unePoint, err := usernameElem.Location()
		if err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to find location of username input element: %v", err)
		}

		// Set mouse movement actions to the username element.
		wd.StorePointerActions("mouse1",
			selenium.MousePointer,
			selenium.PointerMoveAction(0, *unePoint, selenium.FromViewport),
			selenium.PointerPauseAction(100),
			selenium.PointerDownAction(selenium.LeftButton),
			selenium.PointerPauseAction(100),
			selenium.PointerUpAction(selenium.LeftButton),
		)

		// Trigger the mouse movement actions.
		if err := wd.PerformActions(); err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to click username input element: %v", err)
		}

		// Sleep for a bit assuming a human user switches to his keyboard.
		time.Sleep(1*time.Second + time.Duration(r.Intn(100)))

		// Input username to the input tag.
		wd.StoreKeyActions("keyboard1",
			selenium.KeyDownAction("h"),
			selenium.KeyPauseAction(10),
			selenium.KeyDownAction("e"),
			selenium.KeyPauseAction(50),
			selenium.KeyDownAction("l"),
			selenium.KeyPauseAction(20),
			selenium.KeyDownAction("l"),
			selenium.KeyPauseAction(10),
			selenium.KeyDownAction("o"),
		)
		if err := wd.PerformActions(); err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to input username: %v", err)
		}

		// Sleep for a bit assuming the human user switches to his mouse.
		time.Sleep(1*time.Second + time.Duration(r.Intn(100)))

		// Get an input HTML element for password.
		passwordElem, err := wd.FindElement(selenium.ByCSSSelector, "input[name=\"password\"]")
		if err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to find password input element: %v", err)
		}

		// Get the location of the password element.
		pwePoint, err := passwordElem.Location()
		if err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to find location of password input element: %v", err)
		}
		newPoint := selenium.Point{X: pwePoint.X - unePoint.X, Y: pwePoint.Y - unePoint.Y}

		// Set mouse movement actions to the password element.
		wd.StorePointerActions("mouse1",
			selenium.MousePointer,
			selenium.PointerMoveAction(0, newPoint, selenium.FromPointer),
			selenium.PointerPauseAction(250),
			selenium.PointerDownAction(selenium.LeftButton),
			selenium.PointerPauseAction(250),
			selenium.PointerUpAction(selenium.LeftButton),
		)

		// Trigger the mouse movement actions.
		if err := wd.PerformActions(); err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to click password input element: %v", err)
		}

		// Sleep for a bit assuming the human user switches to his keyboard.
		time.Sleep(1*time.Second + time.Duration(r.Intn(100)))

		// Input password to the input tag.
		wd.StoreKeyActions("keyboard1",
			selenium.KeyDownAction("h"),
			selenium.KeyPauseAction(10),
			selenium.KeyDownAction("e"),
			selenium.KeyPauseAction(50),
			selenium.KeyDownAction("l"),
			selenium.KeyPauseAction(10),
			selenium.KeyDownAction("l"),
			selenium.KeyPauseAction(30),
			selenium.KeyDownAction("o"),
		)
		if err := wd.PerformActions(); err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to input password: %v", err)
		}

		// Sleep for a bit assuming the human user switches to his mouse.
		time.Sleep(1*time.Second + time.Duration(r.Intn(100)))

		// Get a checkbox HTML element.
		checkboxElem, err := wd.FindElement(selenium.ByCSSSelector, "input[name=\"checkbox\"]")
		if err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to find checkbox input element: %v", err)
		}

		// Get the location of the checkbox element.
		cbePoint, err := checkboxElem.Location()
		if err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to find location of checkbox input element: %v", err)
		}
		newPoint = selenium.Point{X: cbePoint.X - pwePoint.X, Y: cbePoint.Y - pwePoint.Y}

		// Set mouse movement actions to the checkbox element.
		wd.StorePointerActions("mouse1",
			selenium.MousePointer,
			selenium.PointerMoveAction(0, newPoint, selenium.FromPointer),
			selenium.PointerPauseAction(250),
			selenium.PointerDownAction(selenium.LeftButton),
			selenium.PointerPauseAction(250),
			selenium.PointerUpAction(selenium.LeftButton),
		)

		// Trigger the mouse movement actions.
		if err := wd.PerformActions(); err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to click checkbox element: %v", err)
		}

		// Sleep for a bit.
		time.Sleep(time.Microsecond * 500)

		// Get an HTML login button element.
		loginButtonElem, err := wd.FindElement(selenium.ByCSSSelector, "input[name=\"submit\"]")
		if err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to find login button element: %v", err)
		}

		// Get the location of the login button element.
		lbePoint, err := loginButtonElem.Location()
		if err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to find location of login button element: %v", err)
		}
		newPoint = selenium.Point{X: lbePoint.X - cbePoint.X, Y: lbePoint.Y - cbePoint.Y}

		// Set mouse movement actions to the login button element.
		wd.StorePointerActions("mouse1",
			selenium.MousePointer,
			selenium.PointerMoveAction(0, newPoint, selenium.FromPointer),
			selenium.PointerPauseAction(250),
			selenium.PointerDownAction(selenium.LeftButton),
			selenium.PointerPauseAction(250),
			selenium.PointerUpAction(selenium.LeftButton),
		)

		// Trigger the mouse movement actions.
		if err := wd.PerformActions(); err != nil {
			return fmt.Errorf("[crawler.Run()] Failed to click checkbox element: %v", err)
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
- https://gobyexample.com/random-numbers
- https://deviceatlas.com/blog/list-of-user-agent-strings
- https://www.w3schools.com/browsers/browsers_display.asp
- https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html
- https://chromedriver.chromium.org/capabilities
- https://developer.mozilla.org/en-US/docs/Web/WebDriver/Capabilities/firefoxOptions
- https://github.com/tebeka/selenium/blob/2fb003ac18dced9a6297ce8c03b54de0eddc5fcc/example_test.go#L122-L165
- https://peter.sh/experiments/chromium-command-line-switches/

# Line Count
- Total:      274
- Reused:     0
- Written:    162
- Referenced: 112
*/
