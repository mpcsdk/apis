package chaindata

import (
	v1 "apis/api/chaindata/v1"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) Count(ctx context.Context, req *v1.CountReq) (res *v1.CountRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
