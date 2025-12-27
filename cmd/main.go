package main

import (
	"log"
	"time"

	"linkedin-automation/stealth"
)

func main() {
	// 🔹 Load state
	state, err := stealth.LoadState("state.yaml")
	if err != nil {
		log.Println("No existing state, starting fresh")
		state = &stealth.State{
			DailyLimit: 10, // set your daily connection limit here
		}
	}

	log.Println("🔁 Script run count:", state.RunCount)
	state.RunCount++

	// 🔹 Check daily connection limit
	if state.SentConnections >= state.DailyLimit {
		log.Println("Daily connection limit reached. Stopping.")
		return
	}

	// 🔹 Launch browser
	browser, err := stealth.LaunchBrowser(false)
	if err != nil {
		log.Fatal(err)
	}
	defer browser.MustClose()

	// 🔹 Open LinkedIn feed
	page := browser.MustPage("https://www.linkedin.com/feed/")
	page.MustWaitLoad()

	// Apply fingerprint masking
	stealth.ApplyFingerprintMask(page)

	// 🔹 Human-like scrolling
	log.Println("Scrolling human-like")
	stealth.ScrollHuman(page, 5*time.Second)
	log.Println("Scrolling done")

	// 🔹 Search profiles
	log.Println("Searching profiles")
	profiles, err := stealth.SearchProfiles(page, "Software Engineer", 10)
	if err != nil {
		log.Fatal("Search failed:", err)
	}

	// 🔹 Send connections with daily limit
	count := 0
	for _, profile := range profiles {
		if count >= state.DailyLimit {
			log.Println("Reached daily limit for today")
			break
		}

		log.Println("Sending connection to:", profile)
		err := stealth.SendConnectionRequest(page, profile, "Hi, I’d like to connect with you on LinkedIn!")
		if err != nil {
			log.Println("Failed:", err)
			continue
		}

		state.SentConnections++
		count++
		stealth.SaveState("state.yaml", state)
		time.Sleep(6 * time.Second) // human-like delay
	}

	// 🔹 Send messages to new connections
	connections, err := stealth.GetConnections(page)
	if err != nil {
		log.Println("Failed to get connections:", err)
	} else {
		for _, conn := range connections {
			if conn.ID <= state.SentMessages {
				continue // skip already messaged connections
			}

			log.Println("Sending message to", conn.Name)
			profilePage := browser.MustPage(conn.ProfileURL)
			profilePage.MustWaitLoad()

			err := stealth.SendMessage(profilePage, conn.ProfileURL, "Hi "+conn.Name+", thanks for connecting!")
			if err != nil {
				log.Println("Failed to send message to", conn.Name, ":", err)
			}

			state.SentMessages++
			stealth.SaveState("state.yaml", state)
			time.Sleep(5 * time.Second) // human-like delay
		}
	}

	log.Println("All tasks done!")
}
