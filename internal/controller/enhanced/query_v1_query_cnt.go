package enhanced

import (
	v1 "apis/api/enhanced/v1"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao"
)

func (c *ControllerV1) QueryCnt(ctx context.Context, req *v1.QueryCntReq) (res *v1.QueryCntRes, err error) {
	///
	if req.From == "" && req.ChainId == 0 && req.Contract == "" {
		return nil, mpccode.CodeParamInvalid()
	}

	///
	cnt, err := c.enhanced_riskctrl.GetAggCnt(ctx, mpcdao.QueryEnhancedRiskCtrlRes{
		From:     req.From,
		Contract: req.Contract,
		ChainId:  req.ChainId,
		StartTs:  req.StartTime,
		EndTs:    req.EndTime,
	})
	if err != nil {
		g.Log().Error(ctx, "QueryCnt err:", err)
		return nil, mpccode.CodeInternalError(mpccode.TraceId(ctx))
	}

	return &v1.QueryCntRes{
		Result: cnt,
	}, nil
}
