package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserRoleidPoolModel = (*customUserRoleidPoolModel)(nil)

type (
	// UserRoleidPoolModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRoleidPoolModel.
	UserRoleidPoolModel interface {
		userRoleidPoolModel
	}

	customUserRoleidPoolModel struct {
		*defaultUserRoleidPoolModel
	}
)

// NewUserRoleidPoolModel returns a model for the database table.
func NewUserRoleidPoolModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserRoleidPoolModel {
	return &customUserRoleidPoolModel{
		defaultUserRoleidPoolModel: newUserRoleidPoolModel(conn, c, opts...),
	}
}
