// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package chaindata

import (
	"apis/api/chaindata"
	"apis/internal/service"

	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
	// "apis/internal/service"
	// "github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
)

type ControllerV1 struct{
	contracts map[string]*entity.RiskadminContractabi
	chains map[int64]*entity.RiskadminChaincfg
}


func NewV1() chaindata.IChaindataV1{
	s := &ControllerV1{
		contracts: make(map[string]*entity.RiskadminContractabi),
		chains: make(map[int64]*entity.RiskadminChaincfg),

	}
	////
	// ctx := gctx.GetInitCtx()
	contracts := service.RiskAdmin().RiskAdminCfg().AllContract()
	// if err != nil {
	// 	panic(err)
	// }
	for _, c := range contracts {
		s.contracts[c.ContractAddress] = c
	}
	////
	chains  := service.RiskAdmin().RiskAdminCfg().AllChain()
	for _, c := range chains {
		s.chains[c.ChainId] = c
	}
	////

	return s
}

