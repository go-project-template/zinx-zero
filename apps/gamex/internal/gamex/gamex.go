package gamex

import (
	"sync"
	"zinx-zero/apps/gamex/internal/middleware"
	"zinx-zero/apps/gamex/internal/router"
	"zinx-zero/apps/gamex/internal/svc"

	"github.com/aceld/zinx/zconf"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/service"
)

// GameX is a service for manage zinx server.
type GameX struct {
	service.Service
	Server ziface.IServer
	SvcCtx *svc.ServiceContext
}

var gameXObj *GameX

var onceGameX sync.Once = sync.Once{}

// NewGameX returns a GameX.
func NewGameX(svcCtx *svc.ServiceContext) *GameX {
	onceGameX.Do(func() {
		gameXObj := new(GameX)
		gameXObj.SvcCtx = svcCtx
	})
	return gameXObj
}

// Start runs the server
func (m *GameX) Start() {
	var zinxConf = new(zconf.Config)
	copier.Copy(zinxConf, m.SvcCtx.Config.ZinxConf)

	m.Server = znet.NewUserConfDefaultRouterSlicesServer(zinxConf)
	// 简单累计所有路由组的耗时
	m.Server.Use(znet.RouterTime)
	// deserializing protobuf messages
	m.Server.Use(middleware.RouterProtoUnmarshal)
	// register all handlers
	router.RegisterHandlers(m.Server)
	// runs the server
	m.Server.Serve()
}

// Stop stops the server.
func (m *GameX) Stop() {
	m.Server.Stop()
}
