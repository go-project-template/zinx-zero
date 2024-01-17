package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserAccountModel = (*customUserAccountModel)(nil)

type (
	// UserAccountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAccountModel.
	UserAccountModel interface {
		userAccountModel
	}

	customUserAccountModel struct {
		*defaultUserAccountModel
	}
)

// NewUserAccountModel returns a model for the database table.
func NewUserAccountModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserAccountModel {
	return &customUserAccountModel{
		defaultUserAccountModel: newUserAccountModel(conn, c, opts...),
	}
}
