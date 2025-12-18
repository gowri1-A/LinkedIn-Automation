package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"

	"subspace-automation/stealth"
	"subspace-automation/storage"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("SUBSPACE automation project started!")

	// -----------------------------
	// BUSINESS HOURS CHECK
	// -----------------------------
	if !stealth.IsBusinessHours() {
		fmt.Println("Outside business hours. Exiting like a human.")
		return
	}

	// -----------------------------
	// LAUNCH BROWSER (VISIBLE)
	// -----------------------------
	url := launcher.New().
		Headless(false).
		MustLaunch()

	browser := rod.New().
		ControlURL(url).
		MustConnect()
	defer browser.MustClose()

	// -----------------------------
	// OPEN PAGE + APPLY STEALTH
	// -----------------------------
	page := browser.MustPage()
	stealth.ApplyFingerprintStealth(page)

	// -----------------------------
	// LOAD COOKIES
	// -----------------------------
	fmt.Println("Before loading cookies")
	if _, err := os.Stat("cookies.json"); err == nil {
		fmt.Println("Loading saved cookies...")
		_ = storage.LoadCookies(page, "cookies.json")
	}
	fmt.Println("After loading cookies")

	// -----------------------------
	// NAVIGATE TO LINKEDIN
	// -----------------------------
	fmt.Println("Navigating to LinkedIn feed...")
	page.MustNavigate("https://www.linkedin.com/feed/")
	page.MustWaitLoad()
	time.Sleep(3 * time.Second)

	fmt.Println("Page loaded")
	fmt.Println("Is logged in?", IsLoggedIn(page))

	// -----------------------------
	// MANUAL LOGIN IF REQUIRED
	// -----------------------------
	if !IsLoggedIn(page) {
		fmt.Println("Please login manually, then press ENTER...")
		fmt.Scanln()
		fmt.Println("Saving cookies...")
		_ = storage.SaveCookies(page, "cookies.json")
	}

	// -----------------------------
	// HUMAN-LIKE BEHAVIOR
	// -----------------------------
	fmt.Println("Starting human-like interactions")
	stealth.HumanMouseMovement()
	stealth.HumanScroll(page, 300)
	time.Sleep(2 * time.Second)

	// -----------------------------
	// SEARCH FLOW (STABLE)
	// -----------------------------
	// -----------------------------
	// SEARCH FLOW (STABLE)
	// -----------------------------
	// -----------------------------
	// SEARCH FLOW (KEYBOARD BASED)
	// -----------------------------
	fmt.Println("Typing search query...")

	// Use human-like typing instead of MustInput
	stealth.HumanTyping(page, "input[placeholder='Search']", "harini")
	time.Sleep(1500 * time.Millisecond) // small wait for suggestions to appear

	fmt.Println("Selecting first suggestion using keyboard")

	page.Keyboard.Press(input.ArrowDown)
	time.Sleep(600 * time.Millisecond)

	page.Keyboard.Press(input.Enter)
	page.MustWaitLoad()
	time.Sleep(3 * time.Second)

	fmt.Println("First profile opened successfully")

	time.Sleep(2 * time.Second)

	fmt.Println("Verifying profile page...")
	_, err := page.Element("main")
	if err != nil {
		fmt.Println("Not on profile page. Exiting safely.")
		return
	}

	fmt.Println("Scrolling profile...")
	stealth.HumanScroll(page, 600)
	time.Sleep(2 * time.Second)

	fmt.Println("Looking for Connect button...")

	connectBtn, err := page.ElementR("button", "Connect")
	if err == nil && connectBtn != nil {
		connectBtn.MustScrollIntoView()
		time.Sleep(800 * time.Millisecond)
		connectBtn.MustHover()
		time.Sleep(400 * time.Millisecond)
		connectBtn.MustClick()
		fmt.Println("Clicked Connect button")
	} else {
		fmt.Println("Direct Connect not found, trying More menu...")

		moreBtn, err := page.ElementR("button", "More")
		if err != nil || moreBtn == nil {
			fmt.Println("More button not found. Cannot connect.")
			return
		}

		moreBtn.MustClick()
		time.Sleep(1 * time.Second)

		connectInMore, err := page.ElementR("span", "Connect")
		if err != nil || connectInMore == nil {
			fmt.Println("Connect not available in More menu.")
			return
		}

		connectInMore.MustClick()
		fmt.Println("Clicked Connect from More menu")
	}

	time.Sleep(2 * time.Second)

	sendBtn, err := page.ElementR("button", "Send without a note")
	if err == nil && sendBtn != nil {
		sendBtn.MustClick()
		fmt.Println("Connection request sent")
	} else {
		fmt.Println("Send without note button not found")
	}

	// -----------------------------
	// DONE
	// -----------------------------
	fmt.Println("Automation flow completed.")
	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}

// -----------------------------
// LOGIN CHECK
// -----------------------------
func IsLoggedIn(page *rod.Page) bool {
	_, err := page.Element("img.global-nav__me-photo")
	return err == nil
}
