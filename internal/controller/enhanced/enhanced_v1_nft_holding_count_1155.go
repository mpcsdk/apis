package enhanced

import (
	"context"

	v1 "apis/api/enhanced/v1"
	"apis/internal/service"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao"
	mpcdaoutil "github.com/mpcsdk/mpcCommon/mpcdao/util"
)

func (s *ControllerV1) NftHoldingCount1155(ctx context.Context, req *v1.NftHoldingCount1155Req) (res *v1.NftHoldingCount1155Res, err error) {
	g.Log().Debug(ctx, "NftHoldingCount1155:", "req:", req)
	//specify chainid contractaddr
	if !common.IsHexAddress(req.Address) {
		return nil, mpccode.CodeParamInvalid("address")
	}
	if req.ChainId <= 0 {
		return nil, mpccode.CodeParamInvalid("chainId")
	}
	if !common.IsHexAddress(req.Collection) {
		return nil, mpccode.CodeParamInvalid("collection")
	}
	/////
	if !s.isEnableChain(req.ChainId) {
		return nil, nil
	}
	if !s.isEnableContract(req.ChainId, req.Collection) {
		return nil, nil
	}
	contracts := service.RiskAdmin().RiskAdminCfg().AllContract()
	/////
	////
	rsts, err := s.nftHolding.QueryCount(ctx, &mpcdao.QueryNftHolding{
		ChainId:  req.ChainId,
		Address:  common.HexToAddress(req.Address).String(),
		Contract: common.HexToAddress(req.Collection).String(),
	})
	if err != nil {
		g.Log().Warning(ctx, "NftHoldingCount1155:", "err:", err)
		return nil, mpccode.CodeInternalError()
	}
	g.Log().Debug(ctx, "NftHoldingCount1155:", "rsts:", rsts)
	////
	aggCount := map[string]*v1.NftHolding1155Count{}
	for _, rst := range rsts {
		if abi, ok := contracts[mpcdaoutil.RiskadminContractabiKey(rst.ChainId, rst.Contract)]; !ok {
			g.Log().Info(ctx, "NftHoldingCount1155 not found:", rst.Contract)
			continue
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
