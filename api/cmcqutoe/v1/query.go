package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type QuoteReq struct {
	g.Meta `path:"/quote" tags:"count" method:"post" summary:"You first hello api"`
	Token  string `json:"token"`
}
type QuoteRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Price  float64 `json:"price"`
}

// /
