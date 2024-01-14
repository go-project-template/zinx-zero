package svc

import (
	"zinx-zero/apps/gamex/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config

	RedisClient *redis.Redis

	// UserModel     model.UserModel
	// UserAuthModel model.UserAuthModel
}

// NewServiceContext
func NewServiceContext(c config.Config) *ServiceContext {
	logx.MustSetup(c.Log)
	// sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:      c,
		RedisClient: redis.MustNewRedis(c.Redis.RedisConf),

		// UserAuthModel: model.NewUserAuthModel(sqlConn, c.Cache),
		// UserModel:     model.NewUserModel(sqlConn, c.Cache),
	}
}
