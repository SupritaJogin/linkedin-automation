package stealth

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

// Connection struct
type Connection struct {
	ID         int
	Name       string
	ProfileURL string
}

// GetConnections returns a list of your LinkedIn connections
func GetConnections(page *rod.Page) ([]Connection, error) {
	// Dummy example; replace with actual scraping
	connections := []Connection{
		{ID: 1, Name: "John Doe", ProfileURL: "https://linkedin.com/in/johndoe"},
		{ID: 2, Name: "Jane Smith", ProfileURL: "https://linkedin.com/in/janesmith"},
	}
	return connections, nil
}

// SendMessage sends a message to a connection
func SendMessage(page *rod.Page, profileURL, message string) error {
	profilePage := page.Browser().MustPage(profileURL)
	profilePage.MustWaitLoad()

	// Replace selector with LinkedIn's real message box
	profilePage.MustElement("textarea.msg-form__contenteditable").MustInput(message)
	profilePage.Keyboard.Press(input.Enter) // ✅ use input.Enter instead of "Enter"
	return nil
}
