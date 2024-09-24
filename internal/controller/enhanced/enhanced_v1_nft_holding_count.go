package enhanced

import (
	"context"

	v1 "apis/api/enhanced/v1"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao"
)

func (s *ControllerV1) NftHoldingCount(ctx context.Context, req *v1.NftHoldingCountReq) (*v1.NftHoldingCountRes, error) {
	g.Log().Debug(ctx, "NftHoldingCount:", "req:", req)
	if !common.IsHexAddress(req.Address) {
		return nil, mpccode.CodeParamInvalid("address")
	}
	////
	rsts, err := s.nftHolding.QueryCount(ctx, &mpcdao.QueryNftHolding{
		ChainId: req.ChainId,
		Address: common.HexToAddress(req.Address).String(),
	})
	if err != nil {
		g.Log().Warning(ctx, "NftHoldingCount:", "err:", err)
		return nil, mpccode.CodeInternalError()
	}
	g.Log().Debug(ctx, "NftHoldingCount:", "rsts:", rsts)
	////
	aggCount := map[string]*v1.NftHoldingCount{}
	for _, rst := range rsts {
		if abi, ok := s.contracts[rst.Contract]; !ok {
			return nil, mpccode.CodeParamInvalid(abi.ContractName)
		} else {
			if _, ok := aggCount[abi.ContractName]; !ok {
				aggCount[abi.ContractName] = &v1.NftHoldingCount{
					Symbol:     abi.ContractName,
					Value:      rst.Value,
					Collection: rst.Contract,
				}
			} else {
				aggCount[abi.ContractName].Value += rst.Value
			}
		}
	}
	//////
	res := &v1.NftHoldingCountRes{
		Result: []*v1.NftHoldingCount{},
	}
	for _, v := range aggCount {
		res.Result = append(res.Result, v)
	}
	/////

	return res, nil
}
