package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// TypeHuman types a string with human-like delays
func TypeHuman(page *rod.Page, selector, text string) {
	rand.Seed(time.Now().UnixNano())

	elem := page.MustElement(selector)
	elem.MustClick()

	for _, char := range text {
		// Random delay between 50–250 ms
		delay := time.Duration(rand.Intn(200)+50) * time.Millisecond
		elem.MustInput(string(char))
		time.Sleep(delay)

		// Optional typo simulation (5% chance)
		if rand.Intn(100) < 5 {
			elem.MustInput("\b")        // backspace
			time.Sleep(100 * time.Millisecond)
			elem.MustInput(string(char)) // retype
		}
	}
}
