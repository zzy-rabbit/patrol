package internal

import (
	"errors"
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

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return xerror.ErrAlreadyExists
	}
	return xerror.ErrFail
}
