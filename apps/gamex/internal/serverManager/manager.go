package serverManager

import (
	"zinx-zero/apps/gamex/internal/ice"
	"zinx-zero/apps/gamex/internal/router"
	"zinx-zero/apps/gamex/internal/svc"

	"github.com/aceld/zinx/zconf"
	"github.com/aceld/zinx/zdecoder"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/aceld/zinx/zpack"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/syncx"
)

// Check interface implementation.
var _ ice.IGameServer = (*gameServer)(nil)

var gameServerObj *gameServer

// gameServer is a service for manage zinx server.
type gameServer struct {
	Server ziface.IServer
	SvcCtx *svc.ServiceContext
}

func NewGameServer(svcCtx *svc.ServiceContext) ice.IGameServer {
	syncx.Once(func() {
		gameServerObj = new(gameServer)
		gameServerObj.SvcCtx = svcCtx
	})()
	return gameServerObj
}

func GetGameServer() ice.IGameServer {
	return gameServerObj
}

// Start runs the server
func (m *gameServer) Start() {
	var zinxConf = new(zconf.Config)
	copier.Copy(zinxConf, m.SvcCtx.Config.ZinxConf)

	m.Server = znet.NewUserConfDefaultRouterSlicesServer(zinxConf)

	// Register functions for client connection establishment and loss
	// 注册客户端连接建立和丢失函数
	m.Server.SetOnConnStart(OnConnectionAdd)
	m.Server.SetOnConnStop(OnConnectionLost)

	// 简单累计所有路由组的耗时
	m.Server.Use(znet.RouterTime)
	// deserializing protobuf messages
	// m.Server.Use(middleware.RouterProtoUnmarshal)
	// register all handlers
	router.RegisterHandlers(m.Server)

	// Add LTV data format Decoder
	m.Server.SetDecoder(zdecoder.NewLTV_Little_Decoder())
	// Add LTV data format Pack packet Encoder
	m.Server.SetPacket(zpack.NewDataPackLtv())

	// runs the server
	m.Server.Serve()
}

// Stop stops the server.
func (m *gameServer) Stop() {
	m.Server.Stop()
}

// GetServiceContext implements ice.IGameServer.
func (a *gameServer) GetServiceContext() *svc.ServiceContext {
	return a.SvcCtx
}

func OnConnectionAdd(conn ziface.IConnection) {
	getpl
}
func OnConnectionLost(conn ziface.IConnection) {
}
