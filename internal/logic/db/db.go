package db

import (
	"apis/internal/conf"
	"context"
	"errors"

	"github.com/lib/pq"
	"github.com/mpcsdk/mpcCommon/mpcdao"
	"github.com/mpcsdk/mpcCommon/mpcdao/model/entity"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type sDB struct {
	r             *gredis.Redis
	dur           int
	chainTransfer map[int64]*mpcdao.ChainTransfer
	riskCtrlRule  *mpcdao.RiskCtrlRule
	chainCfg      *mpcdao.ChainCfg
}

func isPgErr(err error, key string) bool {
	gerr := err.(*gerror.Error)
	if cerr, ok := gerr.Cause().(*pq.Error); ok {
		if cerr.Code == pq.ErrorCode(key) {
			return true
		}
	}
	return false
}
func (s *sDB) InitChainTransferDB(ctx context.Context, chainId int64) error {
	err := mpcdao.CreateChainTransferDB(ctx, chainId)
	if err != nil {
		if isPgErr(err, "42P04") {
			///exists
		} else {
			return err
		}
	}
	chaindb := mpcdao.NewChainTransfer(chainId, s.r, s.dur)
	s.chainTransfer[chainId] = chaindb
	return nil
}
func (s *sDB) QueryTransfer(ctx context.Context, chainId int64, query *mpcdao.QueryData) ([]*entity.ChainTransfer, error) {
	// return s.chainTransfer.Query(ctx, query)
	if chaindb, ok := s.chainTransfer[chainId]; ok {
		return chaindb.Query(ctx, query)
	} else {
		g.Log().Error(ctx, "QueryTransfer:", "chainId:", chainId, "query:", query)
		return nil, nil
	}
}

// /
func isDuplicateKeyErr(err error) bool {
	gerr := err.(*gerror.Error)
	if cerr, ok := gerr.Cause().(*pq.Error); ok {
		if cerr.Code == "23505" {
			return true
		}
	}
	return false
}
func (s *sDB) InsertTransfer(ctx context.Context, chainId int64, data *entity.ChainTransfer) error {
	// err := s.chainTransfer.Insert(ctx, data)
	chaindb := s.chainTransfer[chainId]
	if chaindb == nil {
		return errors.New("no chaindb")
	}

	err := chaindb.Insert(ctx, data)
	if err != nil {
		if !isDuplicateKeyErr(err) {
			return err
		}
	}

	return nil
}
func (s *sDB) DelChainBlock(ctx context.Context, chainId int64, block int64) error {
	chaindb := s.chainTransfer[chainId]
	if chaindb == nil {
		return errors.New("no chaindb")
	}
	err := chaindb.DelChainBlockNumber(ctx, chainId, block)
	return err

}
func (s *sDB) InsertTransferBatch(ctx context.Context, chainId int64, datas []*entity.ChainTransfer) error {
	// err := s.chainTransfer.InsertBatch(ctx, datas)
	chaindb := s.chainTransfer[chainId]
	if chaindb == nil {
		return errors.New("no chaindb")
	}
	err := chaindb.InsertBatch(ctx, datas)
	if err != nil {
		return err
	}
	///
	return nil
}

func (s *sDB) ContractAbi() *mpcdao.RiskCtrlRule {
	return s.riskCtrlRule
}
func (s *sDB) ChainCfg() *mpcdao.ChainCfg {
	return s.chainCfg
}
func New() *sDB {

	///
	r := g.Redis()
	_, err := r.Conn(gctx.GetInitCtx())
	if err != nil {
		panic(err)
	}

	///
	s := &sDB{
		r:             r,
		dur:           conf.Config.Cache.Duration,
		chainTransfer: map[int64]*mpcdao.ChainTransfer{},
		//mapmpcdao.NewChainTransfer(r, conf.Config.Cache.SessionDuration),
		riskCtrlRule: mpcdao.NewRiskCtrlRule(r, conf.Config.Cache.Duration),
		chainCfg:     mpcdao.NewChainCfg(r, conf.Config.Cache.Duration),
	}
	////chains cfg
	cfgs, err := s.chainCfg.AllCfg(gctx.GetInitCtx())
	if err != nil {
		panic(err)
	}
	for _, cfg := range cfgs {
		err = s.InitChainTransferDB(gctx.GetInitCtx(), cfg.ChainId)
		if err != nil {
			panic(err)
		}
	}
	return s
}
func init() {

}
