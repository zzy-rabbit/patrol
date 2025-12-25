package internal

import (
	"context"
	"github.com/zzy-rabbit/patrol/data/dao/api"
)

func (s *service) GetDB(ctx context.Context) api.ISession {
	return nil
}

func (s *service) GetTransaction(ctx context.Context) api.ITransaction {
	return nil
}
