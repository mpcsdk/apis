package chaindata

import (
	"context"

	v1 "apis/api/chaindata/v1"
	"apis/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Contract(ctx context.Context, req *v1.ContractReq) (res *v1.ContractRes, err error) {
	g.Log().Debug(ctx, "Query req:", req)
	///
	///
	contracts := service.RiskAdmin().RiskAdminCfg().AllContract()

	//
	res = &v1.ContractRes{
		Contracts: []*v1.ContractResData{},
	}
	for _, contract := range contracts {
		if contract.ChainId != req.ChainId {
			continue
		}
		res.Contracts = append(res.Contracts, &v1.ContractResData{
			ChainId:  contract.ChainId,
			Contract: contract.ContractAddress,
			Name:     contract.ContractName,
			Kind:     contract.ContractKind,
			Decimal:  contract.Decimal,
		})
	}
	return res, nil
}
