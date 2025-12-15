package stealth

import "github.com/go-rod/rod"

func MaskWebDriver(page *rod.Page) {
	page.MustEvalOnNewDocument(`
		Object.defineProperty(navigator, 'webdriver', {
			get: () => undefined
		});
	`)
}
