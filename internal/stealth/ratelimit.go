package stealth

import "time"

type RateLimiter struct {
	Limit int
	Used  int
}

func (r *RateLimiter) Allow() bool {
	if r.Used >= r.Limit {
		time.Sleep(2 * time.Hour) // cooldown
		return false
	}
	r.Used++
	return true
}
