package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RedisCacheKeyModel = (*customRedisCacheKeyModel)(nil)

type (
	// RedisCacheKeyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRedisCacheKeyModel.
	RedisCacheKeyModel interface {
		redisCacheKeyModel
		withSession(session sqlx.Session) RedisCacheKeyModel
	}

	customRedisCacheKeyModel struct {
		*defaultRedisCacheKeyModel
	}
)

// NewRedisCacheKeyModel returns a model for the database table.
func NewRedisCacheKeyModel(conn sqlx.SqlConn) RedisCacheKeyModel {
	return &customRedisCacheKeyModel{
		defaultRedisCacheKeyModel: newRedisCacheKeyModel(conn),
	}
}

func (m *customRedisCacheKeyModel) withSession(session sqlx.Session) RedisCacheKeyModel {
	return NewRedisCacheKeyModel(sqlx.NewSqlConnFromSession(session))
}
