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

func (s *sRateLimiter) Limit() int64 {
	return int64(s.rateLimit.Limit())
}
func (s *sRateLimiter) Tokens() int64 {
	return int64(s.rateLimit.Tokens())
}
