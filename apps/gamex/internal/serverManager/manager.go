package serverManager

import (
	"zinx-zero/apps/gamex/internal/ice"
	"zinx-zero/apps/gamex/internal/middleware"
	"zinx-zero/apps/gamex/internal/router"
	"zinx-zero/apps/gamex/internal/svc"

	"github.com/aceld/zinx/zconf"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/syncx"
)

// Check interface implementation.
var _ ice.IGameServer = (*GameServer)(nil)

var gameServerObj *GameServer

func NewGameServer(svcCtx *svc.ServiceContext) ice.IGameServer {
	syncx.Once(func() {
		gameServerObj = new(GameServer)
		gameServerObj.SvcCtx = svcCtx
	})()
	return gameServerObj
}

func GetGameServer() ice.IGameServer {
	return gameServerObj
}

// GameServer is a service for manage zinx server.
type GameServer struct {
	Server ziface.IServer
	SvcCtx *svc.ServiceContext
}

// Start runs the server
func (m *GameServer) Start() {
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
func (m *GameServer) Stop() {
	m.Server.Stop()
}

// GetServiceContext implements ice.IGameServer.
func (a *GameServer) GetServiceContext() *svc.ServiceContext {
	return a.SvcCtx
}
