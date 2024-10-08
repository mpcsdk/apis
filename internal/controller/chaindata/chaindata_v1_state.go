package chaindata

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"apis/api/chaindata/v1"
)

func (c *ControllerV1) State(ctx context.Context, req *v1.StateReq) (res *v1.StateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
