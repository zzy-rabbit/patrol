package internal

import (
	"context"
	"github.com/zzy-rabbit/patrol/model"
	"github.com/zzy-rabbit/xtools/xerror"
)

func (s *service) SyncDo(ctx context.Context, param model.ExecutorParams) (model.ExecuteResult, xerror.IError) {
	judge := s.NewJudge(ctx, param)
	judge.Exec(ctx)
	return judge.result, nil
}

func (s *service) ASyncDo(ctx context.Context, param model.ExecutorParams, after func(model.ExecuteResult, xerror.IError)) xerror.IError {
	_, err := s.IThreadPool.Do(ctx, func() {
		result, err := s.SyncDo(ctx, param)
		if xerror.Error(err) {
			s.ILogger.Error(ctx, "executor execute task %+v fail %v", param, err)
		}
		after(result, err)
	})
	return err
}
