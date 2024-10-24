package repoimpl

import (
	"context"
	"time"

	"github.com/aide-family/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/pkg/merr"
)

// NewLockRepository 创建全局锁
func NewLockRepository(data *data.Data) repository.Lock {
	return &lockRepositoryImpl{data: data}
}

type lockRepositoryImpl struct {
	data *data.Data
}

func (l *lockRepositoryImpl) Lock(ctx context.Context, key string, expire time.Duration) error {
	exist, err := l.data.GetCacher().Exist(ctx, key)
	if err != nil {
		return err
	}
	// 判断是否存在
	if exist {
		return merr.ErrorI18nToastDatasourceSyncing(ctx)
	}
	return l.data.GetCacher().Set(ctx, key, key, expire)
}

func (l *lockRepositoryImpl) UnLock(ctx context.Context, key string) error {
	return l.data.GetCacher().Delete(ctx, key)
}
