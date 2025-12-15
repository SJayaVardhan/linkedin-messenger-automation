package storage

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

const (
	cookieDir  = "cookies"
	cookieFile = "cookies/session.json"
)

// SaveCookies persists current browser cookies to disk
func SaveCookies(page *rod.Page) error {
	cookies, err := page.Cookies(nil)
	if err != nil {
		return err
	}

	if len(cookies) == 0 {
		return errors.New("no cookies found to save")
	}

	data, err := json.MarshalIndent(cookies, "", "  ")
	if err != nil {
		return err
	}

	if err := os.MkdirAll(cookieDir, 0755); err != nil {
		return err
	}

	return os.WriteFile(cookieFile, data, 0644)
}

// LoadCookies restores valid LinkedIn cookies into the browser session
func LoadCookies(page *rod.Page) error {
	data, err := os.ReadFile(cookieFile)
	if err != nil {
		return err
	}

	var stored []*proto.NetworkCookie
	if err := json.Unmarshal(data, &stored); err != nil {
		return err
	}

	now := float64(time.Now().Unix())
	params := make([]*proto.NetworkCookieParam, 0, len(stored))

	for _, c := range stored {
		// Skip expired cookies (Expires == 0 means session cookie)
		if c.Expires > 0 && float64(c.Expires) < now {
			continue
		}

		// Restore only LinkedIn cookies
		if !isLinkedInDomain(c.Domain) {
			continue
		}

		params = append(params, &proto.NetworkCookieParam{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Expires:  c.Expires,
			HTTPOnly: c.HTTPOnly,
			Secure:   c.Secure,
			SameSite: c.SameSite,
		})
	}

	if len(params) == 0 {
		return errors.New("no valid cookies to restore")
	}

	return page.SetCookies(params)
}

// CookiesExist checks if a saved session file exists
func CookiesExist() bool {
	_, err := os.Stat(cookieFile)
	return err == nil
}

// isLinkedInDomain ensures only LinkedIn cookies are restored
func isLinkedInDomain(domain string) bool {
	return domain == ".linkedin.com" ||
		domain == "www.linkedin.com" ||
		domain == "linkedin.com"
}
