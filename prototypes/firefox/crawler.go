package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tebeka/selenium"
)

const (
	targetURL       = "https://recaptcha-demo.appspot.com/recaptcha-v3-request-scores.php"
	seleniumPath    = "/github.com/tebeka/selenium/vendor/selenium-server.jar"
	geckoDriverPath = "/github.com/tebeka/selenium/vendor/geckodriver"
	port            = 8080
)

func main() {
	log.Println("Start")

	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		log.Fatalln("Set GOPATH environment variable")
	}

	opts := []selenium.ServiceOption{
		// selenium.StartFrameBuffer(),                    // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(goPath + geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),                     // Output debug information to STDERR.
	}

	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(goPath+geckoDriverPath, port, opts...)
	// // service, err := selenium.NewSeleniumService(goPath+seleniumPath, port, opts...)
	if err != nil {
		log.Fatalln("Failed to init selenium: ", err)
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	wd, err := selenium.NewRemote(selenium.Capabilities{"browserName": "firefox"}, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", port))
	if err != nil {
		log.Fatalln("Failed to connect to webDriver: ", err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get(targetURL); err != nil {
		log.Fatalln("Failed to connect to target URL: ", err)
	}

	fmt.Println("End")
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
