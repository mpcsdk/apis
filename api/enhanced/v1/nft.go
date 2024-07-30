package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type NftHoldingReq struct {
	g.Meta         `path:"/nftHolding" tags:"nftHoldingReq" method:"post" summary:"You first hello api"`
	ChainId        int64  `json:"chainId"`
	Address        string `json:"address"`
	Contract       string `json:"contract"`
	CollectionName string `json:"collectionName"`
	Kind           string `json:"kind"`
	///
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
	//
}
type NftHolding struct {
	ChainId     int64  `json:"chainId"`
	Address     string `json:"address"`
	Contract    string `json:"contract"`
	Symbol      string `json:"symbol"`
	TokenId     string `json:"tokenId"`
	Value       int64  `json:"value"`
	BlockNumber int64  `json:"blockNumber"`
	Kind        string `json:"kind"`
}
type NftHoldingRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Result []*NftHolding `json:"result"`
}

// /////
// ////
type NftHoldingCountReq struct {
	g.Meta  `path:"/nftHoldingCount" tags:"nftHoldingCountReq" method:"post" summary:"You first hello api"`
	ChainId int64  `json:"chainId"`
	Address string `json:"address"`
	//
}
type NftHoldingCount struct {
	Symbol string `json:"symbol"`
	Value  int64  `json:"value"`
}
type NftHoldingCountRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Result []*NftHoldingCount `json:"result"`
}
