package internal

import (
	"context"
	daoApi "github.com/zzy-rabbit/patrol/data/dao/api"
	"path/filepath"
)

func (s *service) New(ctx context.Context, unique string) (daoApi.IDatabase, error) {
	database, ok := s.Get(ctx, unique)
	if ok {
		return database, nil
	}

	config := daoApi.Config{
		Driver: "sqlite",
		Sqlite: daoApi.Sqlite{
			File:     filepath.Join(s.config.Path, unique+".db"),
			User:     "",
			Password: "",
		},
	}
	database, err := s.IDao.OpenDatabase(ctx, config)
	if err != nil {
		s.ILogger.Error(ctx, "open database by config %+v fail %v", config, err)
		return nil, err
	}
	s.mutex.Lock()
	s.databaseMap[unique] = database
	s.mutex.Unlock()
	return database, nil
}

func (s *service) get(ctx context.Context, unique string) (daoApi.IDatabase, bool) {
	database, ok := s.databaseMap[unique]
	return database, ok
}

func (s *service) Get(ctx context.Context, unique string) (daoApi.IDatabase, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.get(ctx, unique)
}

func (s *service) Delete(ctx context.Context, unique string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	database, ok := s.get(ctx, unique)
	if !ok {
		return nil
	}
	delete(s.databaseMap, unique)
	go func() {
		err := database.Close(ctx)
		if err != nil {
			s.ILogger.Error(ctx, "close database %+v fail %v", unique, err)
			return
		}
	}()
	return nil
}
