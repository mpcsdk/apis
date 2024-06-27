package enhanced

import (
	v1 "apis/api/enhanced/v1"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/mpcdao"
)

func (c *ControllerV1) Query(ctx context.Context, req *v1.QueryReq) (res *v1.QueryRes, err error) {
	g.Log().Debug(ctx, "Query req:", req)
	if req.From == "" && req.To == "" && req.Contract == "" {
		return nil, mpccode.CodeParamInvalid("from, to, contract can't be all empty")
	}
	if req.StartTime >= req.EndTime {
		return nil, mpccode.CodeParamInvalid("startTime >= endTime")
	}
	if req.Page < 0 || req.PageSize < 0 {
		return nil, mpccode.CodeParamInvalid("page or pageSize invalid")
	}

	///
	query := &mpcdao.QueryTx{
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
	/////query kind
	if req.Kind == "token" {
		query.Kinds = []string{"erc20", "external"}
	} else if req.Kind == "nft" {
		query.Kinds = []string{"erc721", "erc1155"}
	} else {
		query.Kinds = []string{"external", "erc20", "erc721", "erc1155"}
	}
	////
	result, err := c.enhanced_riskctrl.Query(ctx, query)
	if err != nil {
		g.Log().Error(ctx, "Query err:", err)
		return nil, mpccode.CodeParamInvalid(mpccode.TraceId(ctx))
	}
	//
	res = &v1.QueryRes{}
	res.Result = result
	return res, nil
}
