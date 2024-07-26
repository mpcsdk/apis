// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package enhanced

import (
	"apis/api/enhanced"
	"apis/internal/conf"
	"apis/internal/service"
	"math/big"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/mpcdao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

var bigZero = big.NewInt(0)

type ControllerV1 struct {
	redis           *gredis.Redis
	contracts       map[string]*entity.Contractabi
	collectionNames map[string][]*entity.Contractabi
	//db
	enhanced_riskctrl *mpcdao.EnhancedRiskCtrl
	nftHolding        *mpcdao.NftHolding
}

func NewV1() enhanced.IEnhancedV1 {
	s := &ControllerV1{
		contracts:       make(map[string]*entity.Contractabi),
		collectionNames: map[string][]*entity.Contractabi{},
	}
	///
	ctx := gctx.GetInitCtx()
	contracts, err := service.DB().ContractAbi().GetContractAbiBriefs(ctx, 0, "")
	if err != nil {
		panic(err)
	}
	for _, c := range contracts {
		s.contracts[c.ContractAddress] = c
		if _, ok := s.collectionNames[c.ContractName]; ok {
			s.collectionNames[c.ContractName] = append(s.collectionNames[c.ContractName], c)
		} else {
			s.collectionNames[c.ContractName] = []*entity.Contractabi{c}
		}
	}
	///
	r := g.Redis("aggTx")
	_, err = r.Conn(ctx)
	if err != nil {
		panic(err)
	}
	/////
	s.enhanced_riskctrl = mpcdao.NewEnhancedRiskCtrl(r, conf.Config.Cache.Duration)
	s.nftHolding = mpcdao.NewNftHolding()
	s.redis = r

	return s
}
