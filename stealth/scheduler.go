package stealth

import (
	"log"
	"math/rand"
	"time"
)

// IsBusinessHours checks if current time is within working hours
func IsBusinessHours() bool {
	now := time.Now()
	hour := now.Hour()

	// Business hours: 9 AM – 6 PM
	return hour >= 9 && hour <= 18
}

// RandomPause simulates human thinking / breaks
func RandomPause(minSec, maxSec int) {
	sec := rand.Intn(maxSec-minSec+1) + minSec
	log.Printf("Pausing for %d seconds (human break)\n", sec)
	time.Sleep(time.Duration(sec) * time.Second)
}
