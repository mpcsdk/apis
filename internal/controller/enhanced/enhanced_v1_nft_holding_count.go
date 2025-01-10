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

func (s *ControllerV1) NftHoldingCount(ctx context.Context, req *v1.NftHoldingCountReq) (*v1.NftHoldingCountRes, error) {
	g.Log().Debug(ctx, "NftHoldingCount:", "req:", req)
	if !common.IsHexAddress(req.Address) {
		return nil, mpccode.CodeParamInvalid("address")
	}
	///

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
	contracts := service.RiskAdmin().RiskAdminCfg().AllContract()
	aggCount := map[string]*v1.NftHoldingCount{}
	for _, rst := range rsts {
		if !s.isEnableChain(rst.ChainId) {
			g.Log().Warning(ctx, "NftHoldingCount chainId not enable:", rst.ChainId)
			continue
		}
		if contract, ok := contracts[mpcdaoutil.RiskadminContractabiKey(rst.ChainId, rst.Contract)]; !ok {
			g.Log().Info(ctx, "NftHoldingCount not found:", rst.Contract)
			continue
		} else {
			if _, ok := aggCount[contract.ContractName]; !ok {
				aggCount[contract.ContractName] = &v1.NftHoldingCount{
					Symbol: contract.ContractName,
					Value:  rst.Value,
				}
			} else {
				aggCount[contract.ContractName].Value += rst.Value
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
