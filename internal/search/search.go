package search

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-rod/rod"

	"github.com/jayavardhan/linkedin-automation-poc/internal/logger"
	"github.com/jayavardhan/linkedin-automation-poc/internal/stealth"
)

// SearchPeople searches LinkedIn people and returns profile URLs
func SearchPeople(page *rod.Page, query string, maxProfiles int) ([]string, error) {
	searchURL := fmt.Sprintf(
		"https://www.linkedin.com/search/results/people/?keywords=%s",
		query,
	)

	logger.Log.Infof("Searching people with query: %s", query)
	page.MustNavigate(searchURL)
	time.Sleep(5 * time.Second)

	profiles := make(map[string]bool)
	results := []string{}

	for len(results) < maxProfiles {
		// Human behavior
		stealth.HumanScroll(page)
		stealth.HumanDelay(1200, 2500)

		// Find all anchor tags
		links := page.MustElements("a[href]")

		for _, link := range links {
			href, err := link.Attribute("href")
			if err != nil || href == nil {
				continue
			}

			url := *href

			// Filter LinkedIn profile URLs
			if strings.Contains(url, "/in/") && strings.HasPrefix(url, "https://") {
				// Clean tracking params
				if idx := strings.Index(url, "?"); idx != -1 {
					url = url[:idx]
				}

				if !profiles[url] {
					profiles[url] = true
					results = append(results, url)
					logger.Log.Infof("Found profile: %s", url)
				}
			}

			if len(results) >= maxProfiles {
				break
			}
		}

		// Try pagination
		nextBtn, err := page.Element(`button[aria-label="Next"]`)
		if err != nil {
			break
		}

		stealth.HumanDelay(1500, 3000)
		nextBtn.MustClick()
		time.Sleep(5 * time.Second)
	}

	logger.Log.Infof("Collected %d profiles", len(results))
	return results, nil
}
