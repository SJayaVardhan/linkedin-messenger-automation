package auth

import (
	"errors"
	"os"
	"time"

	"github.com/go-rod/rod"

	"github.com/jayavardhan/linkedin-automation-poc/internal/logger"
	"github.com/jayavardhan/linkedin-automation-poc/internal/stealth"
)

func IsLoggedIn(page *rod.Page) bool {
	page.MustNavigate("https://www.linkedin.com/feed/")
	time.Sleep(3 * time.Second)
	return page.MustHas(".global-nav")
}

func Login(page *rod.Page) error {
	email := os.Getenv("LINKEDIN_EMAIL")
	password := os.Getenv("LINKEDIN_PASSWORD")

	if email == "" || password == "" {
		return errors.New("missing LINKEDIN_EMAIL or LINKEDIN_PASSWORD")
	}

	logger.Log.Info("Attempting LinkedIn login")

	page.MustNavigate("https://www.linkedin.com/login")
	time.Sleep(3 * time.Second)

	// Type email
	page.MustElement(`#username`).MustFocus()
	stealth.HumanDelay(800, 1500)
	stealth.HumanType(page, email)

	// Type password
	page.MustElement(`#password`).MustFocus()
	stealth.HumanDelay(800, 1500)
	stealth.HumanType(page, password)

	stealth.HumanDelay(1000, 2000)

	// Click submit
	page.MustElement(`button[type="submit"]`).MustClick()

	time.Sleep(5 * time.Second)

	// Detect captcha or checkpoint
	if page.MustHas(`input[name="challengeId"]`) ||
		page.MustHas(`iframe[src*="captcha"]`) {
		return errors.New("security checkpoint detected (captcha or 2FA). Manual intervention required")

	}

	// Verify login success
	if !IsLoggedIn(page) {
		return errors.New("login failed")
	}

	logger.Log.Info("Login successful")
	return nil
}
