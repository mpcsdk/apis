// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package chaindata

import (
	"context"

	"apis/api/chaindata/v1"
)

type IChaindataV1 interface {
	Count(ctx context.Context, req *v1.CountReq) (res *v1.CountRes, err error)
	Query(ctx context.Context, req *v1.QueryReq) (res *v1.QueryRes, err error)
	State(ctx context.Context, req *v1.StateReq) (res *v1.StateRes, err error)
	Contract(ctx context.Context, req *v1.ContractReq) (res *v1.ContractRes, err error)
}
