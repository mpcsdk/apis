package riskadmin

import (
	"apis/internal/conf"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/riskAdminService/riskAdminServiceNats"
)

type sRiskAdmin struct {
	riskAdminService *riskAdminServiceNats.RiskAdminNatsService

	chainIdLock      sync.RWMutex
	chainIdEnableMap map[int64]int
}

func NewRiskAdmin() *sRiskAdmin {
	ctx := gctx.GetInitCtx()
	redis := g.Redis("cache")
	_, err := redis.Conn(ctx)
	if err != nil {
		panic(err)
	}
	s := &sRiskAdmin{
		chainIdEnableMap: make(map[int64]int),
	}

	riskAdminService, err := riskAdminServiceNats.NewRiskAdminNatsService(ctx,
		riskAdminServiceNats.RiskAdminServiceCfgCfgBuilder().
			// WithConsumeChainFn(s.ConsumeChainCfg).
			// WithConsumeContractFn(s.ConsumeContract).
			// WithConsumeRiskRuleFn(s.ConsumeRiskRule).
			// WithConsumeRiskRuleCheckRespFn(s.consumeRiskRuleCheck).
			WithRedis(redis, conf.Config.Cache.Duration).
			WithUrlTimeOut(conf.Config.Nats.NatsUrl, int64(conf.Config.Nats.TimeOut)),
	)
	if err != nil {
		panic(err)
	}
	s.riskAdminService = riskAdminService
	return s
}

func (s *sRiskAdmin) RiskAdminCfg() *riskAdminServiceNats.RiskAdminCfg {
	return s.riskAdminService.RiskAdminCfg()
}
