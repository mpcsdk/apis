package cmcqutoe

import (
	"context"

	v1 "apis/api/cmcqutoe/v1"
	"apis/internal/service"
)

func (c *ControllerV1) Quote(ctx context.Context, req *v1.QuoteReq) (res *v1.QuoteRes, err error) {
	price := service.Cmc().GetPrice(ctx, req.Token)
	return &v1.QuoteRes{
		Price: price,
	}, nil
}
