// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	IRateLimiter interface {
		Allow() bool
	}
)

var (
	localRateLimiter IRateLimiter
)

func RateLimiter() IRateLimiter {
	if localRateLimiter == nil {
		panic("implement not found for interface IRateLimiter, forgot register?")
	}
	return localRateLimiter
}

func RegisterRateLimiter(i IRateLimiter) {
	localRateLimiter = i
}
