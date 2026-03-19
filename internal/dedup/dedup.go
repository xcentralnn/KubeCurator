package dedup

import (
	"sync"
	"time"
)

var (
	cache  = make(map[string]time.Time)
	mu     sync.Mutex
	window = 5 * time.Minute
)

func ShouldSend(key string) bool {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()

	if last, ok := cache[key]; ok {
		if now.Sub(last) < window {
			return false
		}
	}

	cache[key] = now
	return true
}