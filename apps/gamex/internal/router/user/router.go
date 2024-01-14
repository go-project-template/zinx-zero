package user

import (
	"github.com/aceld/zinx/ziface"
)

func RegisterUserHandlers(server ziface.IServer) {
	server.AddRouterSlices(1, LoginHandle)
}
