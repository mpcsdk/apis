package enhanced

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
	mpcdaoutil "github.com/mpcsdk/mpcCommon/mpcdao/util"

	v1 "apis/api/enhanced/v1"
	"apis/internal/service"
)

func (s *ControllerV1) nftHoldingCollectionName(ctx context.Context, req *v1.NftHoldingReq) (res *v1.NftHoldingRes, err error) {

	if req.Kind == "" {
		return nil, mpccode.CodeParamInvalid("kind")
	}
	if req.CollectionName == "" {
		return nil, mpccode.CodeParamInvalid("conllectionName")
	}
	contracts := service.RiskAdmin().RiskAdminCfg().AllContract()
	////name to contractAddr
	contractAddrs := []string{}
	filteredContract := map[string]*entity.RiskadminContractabi{}
	///build contractAddrs
	for contractKey, contract := range contracts {
		if contract.ContractName == req.CollectionName {
			contractAddrs = append(contractAddrs, contract.ContractAddress)
			filteredContract[contractKey] = contract
		}
	}
	g.Log().Debug(ctx, "NftHolding match collections:", contractAddrs)
	if len(contractAddrs) == 0 {
		return nil, mpccode.CodeParamInvalid("conllectionName")
	}
	//////
	rst, err := s.nftHolding.Query(ctx, &mpcdao.QueryNftHolding{
		ChainId:   req.ChainId,
		Address:   common.HexToAddress(req.Address).String(),
		Contracts: contractAddrs,
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
		reqchain := service.RiskAdmin().RiskAdminCfg().GetChain(nft.ChainId)
		if reqchain.IsEnable == 0 {
			g.Log().Debug(ctx, "NftHolding Query chain not enable:", nft)
			continue
		}
		nfts = append(nfts, &v1.NftHolding{
			ChainId: nft.ChainId,
			Address: nft.Address,
			Symbol: func() string {
				////
				contract := filteredContract[mpcdaoutil.RiskadminContractabiKey(nft.ChainId, nft.Address)]
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
	contracts := service.RiskAdmin().RiskAdminCfg().AllContract()
	/////
	var collection *entity.RiskadminContractabi = nil
	for key, contract := range contracts {
		queryKey := mpcdaoutil.RiskadminContractabiKey(contract.ChainId, contract.ContractAddress)
		if key == queryKey {
			collection = contract
			break
		}
	}
	if collection == nil {
		g.Log().Error(ctx, "NftHoldingCount not found:", req.Collection)
		return nil, mpccode.CodeParamInvalid("collection")
	}
	////
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
				contract := contracts[mpcdaoutil.RiskadminContractabiKey(nft.ChainId, nft.Address)]
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
