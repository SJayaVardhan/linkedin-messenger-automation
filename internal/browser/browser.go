package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// NewBrowser initializes and returns a Rod browser instance
func NewBrowser(headless bool) (*rod.Browser, error) {
	u := launcher.New().
		Headless(headless).
		MustLaunch()

	browser := rod.New().ControlURL(u)

	if err := browser.Connect(); err != nil {
		return nil, err
	}

	return browser, nil
}
