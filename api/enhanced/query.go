// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package enhanced

import (
	v1 "apis/api/enhanced/v1"
	"context"
)

type IQueryV1 interface {
	QueryCnt(ctx context.Context, req *v1.QueryCntReq) (res *v1.QueryCntRes, err error)
	QuerySum(ctx context.Context, req *v1.QuerySumReq) (res *v1.QuerySumRes, err error)
	Query(ctx context.Context, req *v1.QueryReq) (res *v1.QueryRes, err error)
}


