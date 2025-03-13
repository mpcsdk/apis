package packed

import (
	"apis/internal/conf"
	"apis/internal/logic/cmc"
	"apis/internal/logic/db"
	"apis/internal/logic/ratelimit"
	"apis/internal/logic/riskadmin"
	"apis/internal/service"
)

func init() {
	service.RegisterCmc(cmc.NewCMC(
		conf.Config.CMC.QuotesUrl,
		conf.Config.CMC.SymbolList,
		conf.Config.CMC.ApiKey,
		conf.Config.CMC.CmcInterval,
	))
	service.RegisterRiskAdmin(riskadmin.NewRiskAdmin())
	service.RegisterDB(db.New())
	service.RegisterRateLimiter(ratelimit.NewLimiter(conf.Config.Server.RateLimit))
}
