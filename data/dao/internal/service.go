package internal

import (
	"context"
	"github.com/zzy-rabbit/patrol/data/dao/api"
)

func (s *service) GetDB(ctx context.Context) api.ISession {
	return &session{
		db:     s.db,
		tx:     false,
		logger: s.ILogger,
	}
}

func (s *service) GetTransaction(ctx context.Context) api.ITransaction {
	return s.GetDB(ctx).GetTransaction(ctx)
}
