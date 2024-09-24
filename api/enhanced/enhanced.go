// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package enhanced

import (
	"context"

	"apis/api/enhanced/v1"
)

type IEnhancedV1 interface {
	NftHolding(ctx context.Context, req *v1.NftHoldingReq) (res *v1.NftHoldingRes, err error)
	NftHoldingCount(ctx context.Context, req *v1.NftHoldingCountReq) (res *v1.NftHoldingCountRes, err error)
	NftHoldingCount1155(ctx context.Context, req *v1.NftHoldingCount1155Req) (res *v1.NftHoldingCount1155Res, err error)
	QueryCnt(ctx context.Context, req *v1.QueryCntReq) (res *v1.QueryCntRes, err error)
	QuerySum(ctx context.Context, req *v1.QuerySumReq) (res *v1.QuerySumRes, err error)
	Query(ctx context.Context, req *v1.QueryReq) (res *v1.QueryRes, err error)
}
