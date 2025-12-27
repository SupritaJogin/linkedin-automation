package stealth

import (
	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// LaunchBrowser launches Chrome using rod with leakless disabled
func LaunchBrowser(headless bool) (*rod.Browser, error) {
	log.Println("Launching browser (system Chrome, leakless disabled)")

	l := launcher.New().
		Headless(headless).
		Leakless(false). // important
		NoSandbox(true)

	u := l.MustLaunch()

	browser := rod.New().
		ControlURL(u)

	err := browser.Connect()
	if err != nil {
		return nil, err
	}

	return browser, nil
}

// IsLoggedIn checks if LinkedIn shows the header (simple login check)
func IsLoggedIn(page *rod.Page) bool {
	_, err := page.Element("header") // LinkedIn header appears only when logged in
	return err == nil
}
