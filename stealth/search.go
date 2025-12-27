package stealth

import (
	"time"

	"github.com/go-rod/rod"
)

// SearchProfiles searches LinkedIn for profiles by keyword
func SearchProfiles(page *rod.Page, keyword string, maxResults int) ([]string, error) {
	var profiles []string

	page.MustNavigate("https://www.linkedin.com/search/results/people/?keywords=" + keyword)
	page.MustWaitLoad()
	time.Sleep(3 * time.Second) // wait for results to load

	elements, _ := page.Elements("a.search-result__result-link") // adjust selector if needed
	for i, el := range elements {
		if i >= maxResults {
			break
		}
		href, _ := el.Attribute("href")
		if href != nil {
			profiles = append(profiles, *href)
		}
	}

	return profiles, nil
}
