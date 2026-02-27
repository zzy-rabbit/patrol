package internal

import (
	"context"
	"database/sql"
	logApi "github.com/zzy-rabbit/bp/tool/log/api"
	"github.com/zzy-rabbit/patrol/data/dao/api"
	"github.com/zzy-rabbit/xtools/xerror"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type database struct {
	config   api.Config
	db       *gorm.DB
	database *sql.DB
	ILogger  logApi.IPlugin `xplugin:"bp.tool.log"`
}

func (s *service) OpenDatabase(ctx context.Context, config api.Config) (api.IDatabase, xerror.IError) {
	db, err := gorm.Open(sqlite.Open(config.Sqlite.File), &gorm.Config{})
	if err != nil {
		s.ILogger.Error(ctx, "open database by config %+v fail %v", s.GetName(ctx), config, err)
		return nil, xerror.Extend(xerror.ErrInternalError, "open database fail")
	}
	sqliteDB, err := db.DB()
	if err != nil {
		s.ILogger.Error(ctx, "plugin %s open database by config %+v fail %v", s.GetName(ctx), config, err)
		return nil, xerror.Extend(xerror.ErrInternalError, "get database connection fail")
	}
	s.ILogger.Info(ctx, "open database by config %+v success", config)

	err = db.AutoMigrate(&Department{}, &Point{}, &Router{}, &Plan{})
	if err != nil {
		s.ILogger.Error(ctx, "plugin %s auto migrate fail %v", s.GetName(ctx), err)
		return nil, xerror.Extend(xerror.ErrInternalError, "database auto migrate fail")
	}

	return &database{
		config:   config,
		db:       db,
		database: sqliteDB,
		ILogger:  s.ILogger,
	}, nil
}

func (d *database) GetDB(ctx context.Context) api.ISession {
	return &session{
		db:     d.db,
		tx:     false,
		logger: d.ILogger,
	}
}

func (d *database) GetTransaction(ctx context.Context) api.ITransaction {
	return d.GetDB(ctx).GetTransaction(ctx)
}

func (d *database) Close(ctx context.Context) error {
	d.ILogger.Info(ctx, "close database %+v", d.config)
	return d.database.Close()
}
