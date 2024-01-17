package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserAccountAuthModel = (*customUserAccountAuthModel)(nil)

type (
	// UserAccountAuthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAccountAuthModel.
	UserAccountAuthModel interface {
		userAccountAuthModel
	}

	customUserAccountAuthModel struct {
		*defaultUserAccountAuthModel
	}
)

// NewUserAccountAuthModel returns a model for the database table.
func NewUserAccountAuthModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserAccountAuthModel {
	return &customUserAccountAuthModel{
		defaultUserAccountAuthModel: newUserAccountAuthModel(conn, c, opts...),
	}
}
