package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// ScrollHuman scrolls the page randomly like a human
func ScrollHuman(page *rod.Page, duration time.Duration) {
	rand.Seed(time.Now().UnixNano())
	end := time.Now().Add(duration)

	for time.Now().Before(end) {
		// Random scroll amount between -100 and 300 pixels
		dy := rand.Intn(400) - 100
		page.MustEval(`(dy) => { window.scrollBy(0, dy); }`, dy)

		// Random delay between 200–700 ms
		time.Sleep(time.Duration(rand.Intn(500)+200) * time.Millisecond)
	}
}
