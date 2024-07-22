package enhanced

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao"

	v1 "apis/api/enhanced/v1"
)

func (c *ControllerV1) NftHolding(ctx context.Context, req *v1.NftHoldingReq) (res *v1.NftHoldingRes, err error) {
	g.Log().Debug(ctx, "NftHolding:", "req:", req)
	if req.PageSize == 0 {
		return nil, mpccode.CodeParamInvalid("pageSize")
	}
	if req.ChainId == 0 || !common.IsHexAddress(req.Address) {
		return nil, mpccode.CodeParamInvalid("chainId or address")
	}
	if req.Kind == "" {
		return nil, mpccode.CodeParamInvalid("kind")
	}
	//////
	rst, err := c.nftHolding.Query(ctx, &mpcdao.QueryNftHolding{
		ChainId:  req.ChainId,
		Address:  common.HexToAddress(req.Address).String(),
		Kinds:    []string{req.Kind},
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
	})
	if err != nil {
		g.Log().Warning(ctx, "NftHolding Query:", "req:", req, "err:", err)
		return nil, mpccode.CodeInternalError(mpccode.TraceId(ctx))
	}
	///
	nfts := []*v1.NftHolding{}
	for _, nft := range rst {
		nfts = append(nfts, &v1.NftHolding{
			ChainId: nft.ChainId,
			Address: nft.Address,
			Symbol: func() string {
				////
				contract := c.contracts[nft.Contract]
				if contract != nil {
					return contract.ContractName
				}
				return ""
			}(),
			Contract:    nft.Contract,
			TokenId:     nft.TokenId,
			Value:       nft.Value,
			Kind:        nft.Kind,
			BlockNumber: nft.BlockNumber,
		})

	}
	res = &v1.NftHoldingRes{
		Result: nfts,
	}
	////
	return res, nil
}
