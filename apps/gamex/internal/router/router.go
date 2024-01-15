package router

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zinx_app_demo/mmo_game/api"
)

func RegisterHandlers(server ziface.IServer) {

	// Register routers
	server.AddRouter(2, &api.WorldChatApi{})
	server.AddRouter(3, &api.MoveApi{})

	// user.RegisterUserHandlers(server) // register user handlers
}
