package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func HumanType(page *rod.Page, text string) {
	for _, ch := range text {
		page.Keyboard.Type(input.Key(ch))
		time.Sleep(time.Duration(rand.Intn(120)+60) * time.Millisecond)

		// occasional typo
		if rand.Float64() < 0.03 {
			page.Keyboard.Type(input.Key('x'))
			time.Sleep(80 * time.Millisecond)
			page.Keyboard.Type(input.Backspace)
		}
	}
}
