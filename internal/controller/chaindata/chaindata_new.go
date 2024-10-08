// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package chaindata

import (
	"apis/api/chaindata"
	"apis/internal/service"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type ControllerV1 struct{
	contracts map[string]*entity.Contractabi
	chains map[int64]*entity.Chaincfg
}


func NewV1() chaindata.IChaindataV1{
	s := &ControllerV1{
		contracts: make(map[string]*entity.Contractabi),
		chains: make(map[int64]*entity.Chaincfg),
	}
	////
	ctx := gctx.GetInitCtx()
	contracts, err := service.DB().ContractAbi().GetContractAbiBriefs(ctx, 0, "")
	if err != nil {
		panic(err)
	}
	for _, c := range contracts {
		s.contracts[c.ContractAddress] = c
	}
	////
	chains , err := service.DB().ChainCfg().AllCfg(ctx)
	for _, c := range chains {
		s.chains[c.ChainId] = c
	}
	////
	return s
}

