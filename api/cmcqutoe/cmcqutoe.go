// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package cmcqutoe

import (
	"context"

	"apis/api/cmcqutoe/v1"
)

type ICmcqutoeV1 interface {
	Quote(ctx context.Context, req *v1.QuoteReq) (res *v1.QuoteRes, err error)
}
