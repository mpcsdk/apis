package ratelimit

import (
	"golang.org/x/time/rate"
)

type sRateLimiter struct {
	limit     int
	rateLimit *rate.Limiter
}

func NewLimiter(limit int) *sRateLimiter {
	rateLimit := rate.NewLimiter(rate.Limit(limit), limit)
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
