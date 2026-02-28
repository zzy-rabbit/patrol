package internal

import (
	"context"
	logApi "github.com/zzy-rabbit/bp/tool/log/api"
	daoApi "github.com/zzy-rabbit/patrol/data/dao/api"
	databaseApi "github.com/zzy-rabbit/patrol/data/database/api"
	configApi "github.com/zzy-rabbit/patrol/logic/config/api"
	executorApi "github.com/zzy-rabbit/patrol/logic/executor/api"
	"github.com/zzy-rabbit/patrol/model"
	"github.com/zzy-rabbit/xtools/xerror"
	"runtime/debug"
	"time"
)

type departmentTrigger struct {
	ILogger   logApi.IPlugin      `xplugin:"bp.tool.log"`
	IConfig   configApi.IPlugin   `xplugin:"patrol.logic.config"`
	IExecutor executorApi.IPlugin `xplugin:"patrol.logic.executor"`
	IDatabase databaseApi.IPlugin `xplugin:"patrol.data.database"`

	database   daoApi.IDatabase
	interval   time.Duration
	department string
	cancel     context.CancelFunc
}

func (d *departmentTrigger) init(ctx context.Context) xerror.IError {
	database, ok := d.IDatabase.Get(ctx, d.department)
	if !ok {
		d.ILogger.Error(ctx, "department %s database not found", d.department)
		return xerror.Extend(xerror.ErrNotFound, "department "+d.department)
	}
	d.database = database
	return nil
}

func (d *departmentTrigger) startScanMonitor(ctx context.Context) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				d.ILogger.Error(ctx, "scan monitor panic %v stack %s", err, debug.Stack())
			}
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(d.interval):

			}
		}
	}()
}

func (d *departmentTrigger) prepareExecuteParams(ctx context.Context, plan string) (model.ExecutorParams, xerror.IError) {

	return model.ExecutorParams{}, nil
}

func (d *departmentTrigger) stop(ctx context.Context) {
	d.cancel()
}
