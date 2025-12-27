package stealth

import "github.com/go-rod/rod"

// ApplyFingerprintMask hides automation indicators
func ApplyFingerprintMask(page *rod.Page) {
	page.MustEval(`() => {
		// Hide webdriver flag
		Object.defineProperty(navigator, 'webdriver', {
			get: () => undefined
		});

		// Fake plugins
		Object.defineProperty(navigator, 'plugins', {
			get: () => [1, 2, 3]
		});

		// Fake languages
		Object.defineProperty(navigator, 'languages', {
			get: () => ['en-US', 'en']
		});
	}`)
}
