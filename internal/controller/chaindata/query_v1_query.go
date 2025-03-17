package chaindata

import (
	v1 "apis/api/chaindata/v1"
	"apis/internal/service"
	"context"
	"crypto/sha256"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcconsts"
	"github.com/mpcsdk/mpcCommon/mpcdao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"
	"github.com/shengdoushi/base58"
)

var TronAlphabet = base58.NewAlphabet("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func EVMToTronAddress(evmAddr string) string {
	evmAddr = strings.ToLower(evmAddr)
	addr := strings.TrimPrefix(evmAddr, "0x")

	// 2. 添加 0x41 前缀
	addr = "41" + addr
	addrBytes := common.Hex2Bytes(addr)
	// 3. 计算双重 SHA256 哈希
	hash1 := sha256.Sum256(addrBytes)
	hash2 := sha256.Sum256(hash1[:])
	checksum := hash2[0:4]

	addchecksum := append(addrBytes, checksum...)

	// 10. Base58 编码
	tronAddr := base58.Encode(addchecksum, TronAlphabet)
	return tronAddr
}

func wrapOutTxHash(chainId int64, hash string) string {
	switch chainId {
	case mpcconsts.Tron, mpcconsts.TronShasta, mpcconsts.TronNile:
		return strings.TrimPrefix(hash, "0x")
	}
	return hash
}
func wrapOutAddr(chainId int64, addr string) string {
	switch chainId {
	case mpcconsts.Tron, mpcconsts.TronShasta, mpcconsts.TronNile:
		taddr := EVMToTronAddress(addr)
		return taddr
	}
	return addr
}

func wrapInputAddr(chainId int64, addr string) (string, error) {
	if addr == "" {
		return "", nil
	}
	switch chainId {
	case mpcconsts.Tron, mpcconsts.TronShasta, mpcconsts.TronNile:
		baddr, err := address.Base58ToAddress(addr)
		if err != nil {
			return addr, err
		}
		return common.BytesToAddress(baddr.Bytes()).Hex(), nil
	}
	return addr, nil
}
func (c *ControllerV1) Query(ctx context.Context, req *v1.QueryReq) (res *v1.QueryRes, err error) {
	g.Log().Debug(ctx, "Query req:", req)
	///
	if req.ChainId == 0 {
		return nil, mpccode.CodeParamInvalid("need specify chainId")
	}
	if req.From == "" && req.To == "" && req.Contract == "" {
		return nil, mpccode.CodeParamInvalid("from, to, contract can't be all empty")
	}
	////
	req.From, err = wrapInputAddr(req.ChainId, req.From)
	if err != nil {
		g.Log().Error(ctx, "wrapInputAddr err:", err)
		return nil, mpccode.CodeParamInvalid("from addr invalid")
	}
	req.To, err = wrapInputAddr(req.ChainId, req.To)
	if err != nil {
		g.Log().Error(ctx, "wrapInputAddr err:", err)
		return nil, mpccode.CodeParamInvalid("to addr invalid")
	}
	req.Contract, err = wrapInputAddr(req.ChainId, req.Contract)
	if err != nil {
		g.Log().Error(ctx, "wrapInputAddr err:", err)
		return nil, mpccode.CodeParamInvalid("contract addr invalid")
	}
	////

	/////
	reqchain := service.RiskAdmin().RiskAdminCfg().GetChain(req.ChainId)
	if reqchain == nil || reqchain.IsEnable == 0 {
		g.Log().Warning(ctx, "chainId not enable:", req)
		return nil, nil
	}
	//////
	if req.StartTime >= req.EndTime {
		return nil, mpccode.CodeParamInvalid("startTime >= endTime")
	}
	if req.Page < 0 || req.PageSize < 0 {
		return nil, mpccode.CodeParamInvalid("page or pageSize invalid")
	}
	g.Log().Debug(ctx, "Query req:", req)
	///
	query := &mpcdao.QueryData{
		ChainId: req.ChainId,
		From: func() string {
			if req.From == "" {
				return ""
			} else {
				return common.HexToAddress(req.From).String()
			}
		}(),
		To: func() string {
			if req.To == "" {
				return ""
			} else {
				return common.HexToAddress(req.To).String()
			}
		}(),
		Contract: func() string {
			if req.Contract == "" {
				return ""
			} else {
				return common.HexToAddress(req.Contract).String()
			}
		}(),
		///
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		///
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	/////
	if req.Kind == "token" {
		query.Kinds = []string{"erc20", "external"}
	} else if req.Kind == "nft" {
		query.Kinds = []string{"erc721", "erc1155"}
	} else {
		query.Kinds = []string{"external", "erc20", "erc721", "erc1155"}
	}
	////
	results, err := service.DB().QueryTransfer(ctx, query.ChainId, query)
	if err != nil {
		g.Log().Error(ctx, "Query err:", err)
		return nil, mpccode.CodeInternalError(mpccode.TraceId(ctx))
	}
	/////
	res = &v1.QueryRes{}
	for _, r := range results {
		var contract *entity.RiskadminContractabi = nil
		var chain *entity.RiskadminChaincfg = nil
		if r.Kind == "external" {
			chain = service.RiskAdmin().RiskAdminCfg().GetChain(r.ChainId)
			if chain == nil || chain.IsEnable == 0 {
				g.Log().Debug(ctx, "Query unsuppport chainCoin:", r)
				continue
			}
		} else {
			contract = service.RiskAdmin().RiskAdminCfg().GetContract(r.ChainId, r.Contract)
			if contract == nil {
				g.Log().Debug(ctx, "Query unsuppport contract:", r)
				continue
			}
		}

		res.Result = append(res.Result, &v1.QueryResult{
			ChainId:   r.ChainId,
			BlockHash: r.BlockHash,
			TxHash:    wrapOutTxHash(r.ChainId, r.TxHash),
			Ts:        r.Ts,
			From:      wrapOutAddr(r.ChainId, r.From),
			To:        wrapOutAddr(r.ChainId, r.To),
			Contract:  wrapOutAddr(r.ChainId, r.Contract),
			Kind:      r.Kind,
			Status:    r.Status,
			Symbol: func() string {
				////
				if r.Kind == "external" {
					return chain.Coin
					// chain := c.chains[r.ChainId]
					// if chain != nil {
					// 	return chain.Coin
					// }
				} else {
					return contract.ContractName
				}
			}(),
			Value: func() string {
				if r.Kind == "external" {
					fbalance := big.NewFloat(0)
					fbalance.SetString(r.Value)
					chainCfg := service.RiskAdmin().RiskAdminCfg().GetChain(r.ChainId)
					decimal := 18
					if chainCfg == nil {
						g.Log().Warning(ctx, "chainCfg not found:", r.ChainId)
					} else {
						decimal = chainCfg.Decimal
					}
					fval := fbalance.Quo(fbalance, big.NewFloat(math.Pow10(decimal)))
					s := fval.Text('f', -1)
					return s
				} else if r.Kind == "erc20" {
					contract := c.contracts[r.Contract]
					if contract != nil {
						fbalance := big.NewFloat(0)
						fbalance.SetString(r.Value)
						fval := fbalance.Quo(fbalance, big.NewFloat(math.Pow10(contract.Decimal)))
						return fval.Text('f', -1)
					} else {
						return r.Value
					}
				}
				return r.Value
			}(),
			TokenId: r.TokenId,
		})

	}
	return res, nil
}
