package microserver

import (
	"context"

	strategyapi "github.com/aide-family/moon/api/houyi/strategy"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/microrepository"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/cmd/server/palace/internal/service/builder"
	"github.com/go-kratos/kratos/v2/log"
)

// NewStrategyRepository 创建策略仓库
func NewStrategyRepository(data *data.Data, houyiClient *data.HouYiConn) microrepository.Strategy {
	return &strategyRepositoryImpl{data: data, houyiClient: houyiClient}
}

type strategyRepositoryImpl struct {
	data *data.Data

	houyiClient *data.HouYiConn
}

func (s *strategyRepositoryImpl) Push(ctx context.Context, strategies []*bo.Strategy) error {
	apiStrategies := builder.NewParamsBuild().WithContext(ctx).StrategyModuleBuilder().BoStrategyBuilder().ToAPIs(strategies)
	strategyReply, err := s.houyiClient.PushStrategy(ctx, &strategyapi.PushStrategyRequest{
		Strategies: apiStrategies,
	})
	if err != nil {
		return err
	}
	log.Debugw("strategyReply", strategyReply)
	return nil
}
