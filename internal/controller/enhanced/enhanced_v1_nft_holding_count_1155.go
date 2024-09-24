package enhanced

import (
	"context"

	v1 "apis/api/enhanced/v1"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao"
)

func (s *ControllerV1) NftHoldingCount1155(ctx context.Context, req *v1.NftHoldingCount1155Req) (res *v1.NftHoldingCount1155Res, err error) {
	g.Log().Debug(ctx, "NftHoldingCount1155:", "req:", req)
	if !common.IsHexAddress(req.Address) {
		return nil, mpccode.CodeParamInvalid("address")
	}
	if req.ChainId <= 0 {
		return nil, mpccode.CodeParamInvalid("chainId")
	}
	if !common.IsHexAddress(req.Collection) {
		return nil, mpccode.CodeParamInvalid("collection")
	}
	////
	rsts, err := s.nftHolding.QueryCount(ctx, &mpcdao.QueryNftHolding{
		ChainId:   req.ChainId,
		Address:   common.HexToAddress(req.Address).String(),
		Contracts: []string{common.HexToAddress(req.Collection).String()},
	})
	if err != nil {
		g.Log().Warning(ctx, "NftHoldingCount1155:", "err:", err)
		return nil, mpccode.CodeInternalError()
	}
	g.Log().Debug(ctx, "NftHoldingCount1155:", "rsts:", rsts)
	////
	aggCount := map[string]*v1.NftHolding1155Count{}
	for _, rst := range rsts {
		if abi, ok := s.contracts[rst.Contract]; !ok {
			return nil, mpccode.CodeParamInvalid(abi.ContractName)
		} else {
			if _, ok := aggCount[abi.ContractAddress]; !ok {
				aggCount[abi.ContractAddress] = &v1.NftHolding1155Count{
					Value:      rst.Value,
					Collection: abi.ContractAddress,
				}
			} else {
				aggCount[abi.ContractAddress].Value += rst.Value
			}
		}
	}
	//////
	res = &v1.NftHoldingCount1155Res{
		Result: []*v1.NftHolding1155Count{},
	}
	for _, v := range aggCount {
		res.Result = append(res.Result, v)
	}
	/////

	return res, nil
}
