package stealth

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// -----------------------------
// Business hours check (DEV MODE)
// -----------------------------

// -----------------------------
// Utility: random delay
// -----------------------------
func randomDelay(minMs, maxMs int) {
	time.Sleep(time.Duration(rand.Intn(maxMs-minMs)+minMs) * time.Millisecond)
}

// -----------------------------
// Fingerprint masking
// -----------------------------
func ApplyFingerprintStealth(page *rod.Page) {
	fmt.Println("Applying browser fingerprint stealth")

	js := `
		Object.defineProperty(navigator, 'webdriver', {
			get: () => undefined
		});

		window.chrome = {
			runtime: {}
		};

		Object.defineProperty(navigator, 'languages', {
			get: () => ['en-US', 'en']
		});

		Object.defineProperty(navigator, 'plugins', {
			get: () => [1, 2, 3, 4, 5]
		});
	`
	page.EvalOnNewDocument(js)
}

// -----------------------------
// Mouse movement (mock human)
// -----------------------------
func HumanMouseMovement() {
	fmt.Println("Simulating human mouse movement")
	steps := rand.Intn(5) + 3
	for i := 0; i < steps; i++ {
		randomDelay(200, 500)
	}
}

// -----------------------------
// Human scroll
// -----------------------------
func HumanScroll(page *rod.Page, maxScroll int) {
	fmt.Println("Simulating human scroll")

	scrollSteps := rand.Intn(5) + 3
	for i := 0; i < scrollSteps; i++ {
		scroll := rand.Intn(maxScroll) + 50
		_, _ = page.Eval(`window.scrollBy(0, ` + fmt.Sprint(scroll) + `)`)
		randomDelay(300, 700)
	}

	if rand.Intn(2) == 1 {
		_, _ = page.Eval(`window.scrollBy(0, -` + fmt.Sprint(rand.Intn(120)+30) + `)`)
		randomDelay(200, 400)
	}
}

// -----------------------------
// Human typing with typos (SAFE)
// -----------------------------
func HumanTyping(page *rod.Page, selector string, text string) {
	fmt.Println("Simulating human typing for:", selector)

	el := page.MustElement(selector)
	el.MustClick() // focus on input

	for _, ch := range text {
		el.MustInput(string(ch))
		time.Sleep(time.Duration(rand.Intn(150)+100) * time.Millisecond) // small delay for realism
	}
}

func randomChar() string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	return string(letters[rand.Intn(len(letters))])
}
