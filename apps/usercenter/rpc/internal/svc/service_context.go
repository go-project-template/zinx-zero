package svc

import (
	"zinx-zero/apps/acommon/globalkey"
	"zinx-zero/apps/model"
	"zinx-zero/apps/usercenter/rpc/internal/config"

	"github.com/aceld/zinx/zutils"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	IDWorker *zutils.IDWorker

	RedisClient *redis.Redis

	UserAccountModel     model.UserAccountModel
	UserAccountAuthModel model.UserAccountAuthModel
	UserRoleModel        model.UserRoleModel
	UserRoleidPoolModel  model.UserRoleidPoolModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	idWorker, err := zutils.NewIDWorker(globalkey.SnowflakeWorkerId_Usercenter)
	logx.Must(err)

	return &ServiceContext{
		Config:      c,
		IDWorker:    idWorker,
		RedisClient: redis.MustNewRedis(c.Redis.RedisConf),

		UserAccountModel:     model.NewUserAccountModel(sqlConn, c.Cache),
		UserAccountAuthModel: model.NewUserAccountAuthModel(sqlConn, c.Cache),
		UserRoleModel:        model.NewUserRoleModel(sqlConn, c.Cache),
		UserRoleidPoolModel:  model.NewUserRoleidPoolModel(sqlConn),
	}
}
