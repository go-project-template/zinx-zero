package ice

import (
	"zinx-zero/apps/gamex/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

type (
	IGameServer interface {
		service.Service
		GetServiceContext() *svc.ServiceContext
	}
)
