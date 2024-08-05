package enhanced

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao"

	v1 "apis/api/enhanced/v1"
)

func (s *ControllerV1) nftHoldingCollectionName(ctx context.Context, req *v1.NftHoldingReq) (res *v1.NftHoldingRes, err error) {

	if req.Kind == "" {
		return nil, mpccode.CodeParamInvalid("kind")
	}
	if req.CollectionName == "" {
		return nil, mpccode.CodeParamInvalid("conllectionName")
	}

	////name to contracts
	contracts := []string{}
	if abis, ok := s.collectionNames[req.CollectionName]; !ok {
		return nil, mpccode.CodeParamInvalid(req.CollectionName)
	} else {
		for _, abi := range abis {
			contracts = append(contracts, abi.ContractAddress)
		}
	}
	g.Log().Debug(ctx, "NftHolding match collections:", contracts)
	//////
	rst, err := s.nftHolding.Query(ctx, &mpcdao.QueryNftHolding{
		ChainId:   req.ChainId,
		Address:   common.HexToAddress(req.Address).String(),
		Contracts: contracts,
		Kinds:     []string{req.Kind},
		Page:      int(req.Page),
		PageSize:  int(req.PageSize),
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
				contract := s.contracts[nft.Contract]
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
func (s *ControllerV1) nftHoldingCollection(ctx context.Context, req *v1.NftHoldingReq) (res *v1.NftHoldingRes, err error) {

	if req.ChainId == 0 {
		return nil, mpccode.CodeParamInvalid("chainId")
	}
	if !common.IsHexAddress(req.Collection) {
		return nil, mpccode.CodeParamInvalid("collection")
	}
	g.Log().Debug(ctx, "NftHolding match collections:", req.Collection)
	//////
	rst, err := s.nftHolding.Query(ctx, &mpcdao.QueryNftHolding{
		ChainId:   req.ChainId,
		Address:   common.HexToAddress(req.Address).String(),
		Contracts: []string{common.HexToAddress(req.Collection).String()},
		Page:      int(req.Page),
		PageSize:  int(req.PageSize),
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
				contract := s.contracts[nft.Contract]
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
func (s *ControllerV1) NftHolding(ctx context.Context, req *v1.NftHoldingReq) (res *v1.NftHoldingRes, err error) {
	g.Log().Debug(ctx, "NftHolding:", "req:", req)
	if req.PageSize == 0 {
		return nil, mpccode.CodeParamInvalid("pageSize")
	}
	if !common.IsHexAddress(req.Address) {
		return nil, mpccode.CodeParamInvalid("address")
	}
	////
	if req.Collection != "" {
		return s.nftHoldingCollection(ctx, req)

	} else {
		return s.nftHoldingCollectionName(ctx, req)
	}

}
