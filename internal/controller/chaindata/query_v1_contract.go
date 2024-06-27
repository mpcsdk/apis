package chaindata

import (
	"context"

	v1 "apis/api/chaindata/v1"
	"apis/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (c *ControllerV1) Contract(ctx context.Context, req *v1.ContractReq) (res *v1.ContractRes, err error) {
	g.Log().Debug(ctx, "Query req:", req)
	///
	///
	result, err := service.DB().ContractAbi().GetContractAbiBriefs(ctx, req.ChainId, "")
	if err != nil {
		g.Log().Error(ctx, "Query err:", err)
		return nil, mpccode.CodeInternalError(mpccode.TraceId(ctx))
	}
	//
	res = &v1.ContractRes{
		Contracts: []*v1.ContractResData{},
	}
	for _, r := range result {
		res.Contracts = append(res.Contracts, &v1.ContractResData{
			ChainId:  r.ChainId,
			Contract: r.ContractAddress,
			Name:     r.ContractName,
			Kind:     r.ContractKind,
			Decimal:  r.Decimal,
		})
	}
	return res, nil
}
