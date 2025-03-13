package cmc

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type sCmc struct {
	ctx        context.Context
	baseUrl    string
	quotesUrl  string
	symbolList string
	apiKey     string
	interval   int
	////
	redis *gredis.Redis
}

func NewCMC(quotesUrl, symbolList, apiKey string, interval int) *sCmc {
	ctx := gctx.GetInitCtx()
	r := g.Redis()
	_, err := r.Conn(ctx)
	if err != nil {
		panic(err)
	}
	///

	s := &sCmc{
		ctx:        ctx,
		quotesUrl:  quotesUrl,
		symbolList: symbolList,
		apiKey:     apiKey,
		interval:   interval,
		////
		redis: r,
	}
	go s.run()
	return s

}

// ///////
func (s *sCmc) run() {
	for {
		s.syncPrice()
		time.Sleep(time.Second * time.Duration(s.interval))
	}
}

// ///////
func (s *sCmc) syncPrice() {
	// url := fmt.Sprintf("%s?symbol=%s&convert=USD", s.baseUrl, s.symbolList)
	//apiKey := GetCoinMarketCapAPIKey()
	ok, err := s.redis.SetNX(s.ctx, "CMC:Lock:Price", "1")
	if err != nil {
		g.Log().Error(s.ctx, "syncPrice setnx err:", err)
		return
	}
	if !ok {
		g.Log().Debug(s.ctx, "syncPrice setnx not ok:")
		return
	}
	s.redis.SetEX(s.ctx, "CMC:Lock:Price", "1", int64(s.interval))
	////
	r, err := g.Client().
		SetHeader("Accepts", "application/json").
		SetHeader("X-CMC_PRO_API_KEY", s.apiKey).
		Get(s.ctx, s.quotesUrl, map[string]string{
			"symbol":  s.symbolList,
			"convert": "USD",
		})
	if err != nil {
		g.Log().Error(s.ctx, "syncprice quto error:", err)
		return
	}
	defer r.Close()
	////
	var resp *CMCResp = nil
	err = json.Unmarshal(r.ReadAll(), &resp)
	if err != nil {
		g.Log().Error(s.ctx, "syncprice unmarshal error:", err)
		return
	}
	if resp.Status.ErrorCode != 0 {
		g.Log().Error(s.ctx, "syncprice error:", resp.Status.ErrorMessage)
		return
	}

	for token, tokenData := range resp.Data {
		if token == "TRX" {
			g.Log().Debug(s.ctx, "syncprice TRX:", tokenData.Quote.Usd.Price)
			s.redis.Set(s.ctx, "CMC:Price:TRX", tokenData.Quote.Usd.Price)
		}
	}

}
func (s *sCmc) GetPrice(ctx context.Context, token string) float64 {
	if token == "TRX" {
		v, err := s.redis.Get(s.ctx, "CMC:Price:TRX")
		if err != nil {
			g.Log().Error(s.ctx, "GetPrice error:", err)
			return 0
		}
		return v.Float64()
	}
	return 0
}
