package svc

import (
	"autoGenCache/internal/config"
	"autoGenCache/model"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	RedisCacheKeyModel model.RedisCacheKeyModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	_, err := sqlConn.RawDB()
	logx.Must(err)
	return &ServiceContext{
		Config:             c,
		RedisCacheKeyModel: model.NewRedisCacheKeyModel(sqlConn),
	}
}
