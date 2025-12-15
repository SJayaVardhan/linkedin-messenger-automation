package main

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/jayavardhan/linkedin-automation-poc/internal/auth"
	"github.com/jayavardhan/linkedin-automation-poc/internal/browser"
	"github.com/jayavardhan/linkedin-automation-poc/internal/connect"
	"github.com/jayavardhan/linkedin-automation-poc/internal/logger"
	"github.com/jayavardhan/linkedin-automation-poc/internal/search"
	"github.com/jayavardhan/linkedin-automation-poc/internal/stealth"
	"github.com/jayavardhan/linkedin-automation-poc/internal/storage"
)

func main() {
	// Load environment variables
	_ = godotenv.Load()

	// Initialize logger
	logger.Init()
	logger.Log.Info("Starting LinkedIn Automation PoC")

	// Headless mode
	headless := os.Getenv("HEADLESS") == "true"

	// Launch browser
	br, err := browser.NewBrowser(headless)
	if err != nil {
		logger.Log.Fatal(err)
	}

	// Open LinkedIn homepage
	page := br.MustPage("https://www.linkedin.com")

	// ---------------- STEALTH LAYER ----------------
	stealth.MaskWebDriver(page)
	stealth.HumanDelay(1200, 2500)
	stealth.RandomHover(page)
	stealth.HumanScroll(page)

	logger.Log.Info("Browser launched with stealth behavior applied")

	// ---------------- SESSION HANDLING ----------------
	if storage.CookiesExist() {
		logger.Log.Info("Existing session detected, loading cookies")
		if err := storage.LoadCookies(page); err != nil {
			logger.Log.Warn("Failed to load cookies, continuing without session")
		}
	}

	// ---------------- AUTHENTICATION ----------------
	if !auth.IsLoggedIn(page) {
		logger.Log.Info("Not logged in, starting login flow")

		if err := auth.Login(page); err != nil {
			logger.Log.Warn(err.Error())
			logger.Log.Warn("Manual login required. Complete CAPTCHA/2FA in browser.")

			// Wait for manual login instead of exiting
			for {
				stealth.HumanDelay(3000, 5000)

				if auth.IsLoggedIn(page) {
					logger.Log.Info("Manual login detected")

					if err := storage.SaveCookies(page); err != nil {
						logger.Log.Warn("Failed to save session cookies after manual login")
					} else {
						logger.Log.Info("Session cookies saved successfully")
					}
					break
				}
			}
		} else {
			// Login succeeded automatically
			if err := storage.SaveCookies(page); err != nil {
				logger.Log.Warn("Failed to save session cookies")
			} else {
				logger.Log.Info("Session cookies saved successfully")
			}
		}
	} else {
		logger.Log.Info("Logged in using existing session")
	}

	// ---------------- SEARCH ----------------
	profiles, err := search.SearchPeople(page, "software engineer", 10)
	if err != nil {
		logger.Log.Warn(err)
	}

	logger.Log.Infof("Total profiles collected: %d", len(profiles))

	// ---------------- CONNECT ----------------
	connect.SendConnectionRequests(
		page,
		profiles,
		connect.ConnectOptions{
			MaxPerRun: 2,
			Note:      "Hi! I came across your profile and would love to connect.",
		},
	)

	// Keep browser open for observation
	select {}
}
