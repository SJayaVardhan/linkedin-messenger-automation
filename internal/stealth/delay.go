package stealth

import (
	"math/rand"
	"time"
)

func HumanDelay(minMs, maxMs int) {
	rand.Seed(time.Now().UnixNano())
	delay := rand.Intn(maxMs-minMs+1) + minMs
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
