package connect

import (
	"strings"
	"time"

	"github.com/go-rod/rod"

	"github.com/jayavardhan/linkedin-automation-poc/internal/logger"
	"github.com/jayavardhan/linkedin-automation-poc/internal/stealth"
)

type ConnectOptions struct {
	MaxPerRun int
	Note      string
}

// SendConnectionRequests visits profiles and sends connection requests safely
func SendConnectionRequests(
	page *rod.Page,
	profiles []string,
	opts ConnectOptions,
) {
	sent := 0

	for _, profileURL := range profiles {
		if sent >= opts.MaxPerRun {
			logger.Log.Warn("Daily connect limit reached")
			return
		}

		logger.Log.Info("Attempting to send connection request")
		page.MustNavigate(profileURL)
		time.Sleep(4 * time.Second)

		// Human-like behavior
		stealth.HumanDelay(1500, 3000)
		stealth.RandomHover(page)

		// Collect all buttons on the page
		buttons := page.MustElements("button")
		var connectBtn *rod.Element

		for _, btn := range buttons {
			text, err := btn.Text()
			if err != nil {
				continue
			}

			label := strings.ToLower(strings.TrimSpace(text))

			// Skip already-connected or unavailable states
			if strings.Contains(label, "message") ||
				strings.Contains(label, "pending") ||
				strings.Contains(label, "following") {
				logger.Log.Info("Already connected or request pending, skipping")
				connectBtn = nil
				break
			}

			// Identify Connect button explicitly
			if label == "connect" {
				connectBtn = btn
				break
			}
		}

		if connectBtn == nil {
			logger.Log.Info("Connect button not found (likely Follow / Pending / limit reached)")
			continue
		}

		// Click Connect
		stealth.HumanDelay(800, 1500)
		connectBtn.MustClick()
		time.Sleep(2 * time.Second)

		// Optional personalized note
		if opts.Note != "" {
			noteBtn, err := page.Element(`button[aria-label*="Add a note"]`)
			if err == nil {
				stealth.HumanDelay(500, 1000)
				noteBtn.MustClick()
				time.Sleep(1 * time.Second)

				textarea := page.MustElement("textarea")
				textarea.MustFocus()
				stealth.HumanType(page, opts.Note)

				sendBtn := page.MustElement(`button[aria-label*="Send"]`)
				stealth.HumanDelay(800, 1500)
				sendBtn.MustClick()
			}
		}

		sent++
		logger.Log.Infof("Connection request sent (%d/%d)", sent, opts.MaxPerRun)

		// Cooldown between profiles
		stealth.HumanDelay(5000, 9000)
	}
}
