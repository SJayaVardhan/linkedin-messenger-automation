package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func HumanScroll(page *rod.Page) {
	scrolls := rand.Intn(4) + 3
	for i := 0; i < scrolls; i++ {
		page.Mouse.Scroll(
			0,
			float64(rand.Intn(400)+200),
			1,
		)
		time.Sleep(time.Duration(rand.Intn(700)+300) * time.Millisecond)
	}
}
