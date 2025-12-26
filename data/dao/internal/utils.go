package internal

import (
	"errors"
	"github.com/mattn/go-sqlite3"
	"github.com/zzy-rabbit/xtools/xerror"
	"gorm.io/gorm"
)

func transError(err error) xerror.IError {
	if err == nil {
		return nil
	}

	// 无记录
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return xerror.ErrNotFound
	}

	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		if errors.Is(err, sqlite3.ErrConstraintUnique) {
			return xerror.ErrAlreadyExists
		}
	}
	return xerror.ErrFail
}
