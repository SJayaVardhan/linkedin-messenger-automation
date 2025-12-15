package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func RandomHover(page *rod.Page) {
	x := rand.Intn(800) + 100
	y := rand.Intn(600) + 100

	page.MustEval(`() => {
		document.dispatchEvent(
			new MouseEvent('mousemove', {
				clientX: arguments[0],
				clientY: arguments[1],
				bubbles: true
			})
		)
		return true
	}`, x, y)

	time.Sleep(time.Duration(rand.Intn(800)+400) * time.Millisecond)
}
