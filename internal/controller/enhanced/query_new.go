// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package enhanced

import (
	"apis/api/enhanced"
	"apis/internal/conf"
	"math/big"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/mpcdao"
)
var bigZero = big.NewInt(0)
type ControllerV1 struct{
	redis  *gredis.Redis
	//
	enhanced_riskctrl *mpcdao.EnhancedRiskCtrl
	///
}

func NewV1() enhanced.IQueryV1 {
	///
	///
	r := g.Redis("aggTx")
	_, err := r.Conn(gctx.GetInitCtx())
	if err != nil {
		panic(err)
	}
	return &ControllerV1{
		redis:  r,
		enhanced_riskctrl: mpcdao.NewEnhancedRiskCtrl(r,conf.Config.Cache.Duration),
	}
}

