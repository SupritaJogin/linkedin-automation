package stealth

import (
	"log"
	"time"

	"github.com/go-rod/rod"
)

// SendConnectionRequest visits a profile and sends a connection request
func SendConnectionRequest(page *rod.Page, profileURL string, message string) error {
	page.MustNavigate(profileURL)
	page.MustWaitLoad()
	time.Sleep(2 * time.Second) // human-like pause

	connectBtn, err := page.Element("button[data-control-name='connect']") // LinkedIn Connect button
	if err != nil {
		log.Println("Connect button not found on", profileURL)
		return err
	}

	connectBtn.MustClick()
	time.Sleep(1 * time.Second)

	// Add a note
	noteBtn, err := page.Element("button[aria-label='Add a note']")
	if err == nil {
		noteBtn.MustClick()
		time.Sleep(500 * time.Millisecond)
		page.MustElement("textarea[name='message']").MustInput(message)
		time.Sleep(1 * time.Second)
	}

	// Send request
	sendBtn, err := page.Element("button[aria-label='Send now']")
	if err == nil {
		sendBtn.MustClick()
		log.Println("Connection request sent to:", profileURL)
	} else {
		log.Println("Send button not found, maybe request already sent:", profileURL)
	}

	time.Sleep(3 * time.Second) // human-like pause
	return nil
}
