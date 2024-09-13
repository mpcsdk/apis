package packed

import (
	"apis/internal/conf"
	"apis/internal/logic/db"
	"apis/internal/logic/ratelimit"
	"apis/internal/service"
)

func init() {
	service.RegisterRateLimiter(ratelimit.NewLimiter(conf.Config.Server.RateLimit))
	service.RegisterDB(db.New())
}
