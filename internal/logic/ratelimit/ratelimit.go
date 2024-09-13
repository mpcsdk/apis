package ratelimit

import (
	"time"

	"golang.org/x/time/rate"
)

type sRateLimiter struct {
	limit     int
	rateLimit *rate.Limiter
}

func NewLimiter(limit int) *sRateLimiter {
	every := rate.Every(time.Second)
	rateLimit := rate.NewLimiter(every, limit)
	return &sRateLimiter{
		limit:     limit,
		rateLimit: rateLimit,
	}
}

func (s *sRateLimiter) Allow() bool {
	return s.rateLimit.Allow()
}
