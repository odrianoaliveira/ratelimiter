package ratelimit

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity   int
	tokens     int
	refillRate int
	interval   time.Duration
	mutex      sync.Mutex
}

func NewTokenBucket(capacity int, refillRate int, interval time.Duration) *TokenBucket {
	tb := &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		interval:   interval,
	}

	go tb.refill()

	return tb
}

func (tb *TokenBucket) Allow() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

func (tb *TokenBucket) refill() {
	ticker := time.NewTicker(tb.interval)
	for {
		<-ticker.C
		tb.mutex.Lock()
		tb.tokens += tb.refillRate
		if tb.tokens > tb.capacity {
			tb.tokens = tb.capacity
		}
		tb.mutex.Unlock()
	}
}
