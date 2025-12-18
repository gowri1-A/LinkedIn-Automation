package stealth

import (
	"fmt"
	"math/rand"
	"time"
)

// Allow activity only during business hours
func IsBusinessHours() bool {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	now := time.Now().In(loc)

	// Block weekends
	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		return false
	}

	// Business hours: 9:30 AM – 6:30 PM
	start := time.Date(
		now.Year(), now.Month(), now.Day(),
		9, 30, 0, 0, loc,
	)

	end := time.Date(
		now.Year(), now.Month(), now.Day(),
		18, 30, 0, 0, loc,
	)

	return now.After(start) && now.Before(end)
}

// Simulate human break
func TakeBreak(minSec, maxSec int) {
	if maxSec <= minSec {
		maxSec = minSec + 1
	}

	delaySec := rand.Intn(maxSec-minSec) + minSec
	delay := time.Duration(delaySec) * time.Second

	fmt.Println("Taking human break for", delay)
	time.Sleep(delay)
}
