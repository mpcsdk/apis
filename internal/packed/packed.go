package packed

import (
	"apis/internal/conf"
	"apis/internal/logic/db"
	"apis/internal/logic/ratelimit"
	"apis/internal/logic/riskadmin"
	"apis/internal/service"
)

func init() {
	service.RegisterRiskAdmin(riskadmin.NewRiskAdmin())
	service.RegisterDB(db.New())
	service.RegisterRateLimiter(ratelimit.NewLimiter(conf.Config.Server.RateLimit))
}
