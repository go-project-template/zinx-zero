package router

import (
	"zinx-zero/apps/gamex/internal/router/user"

	"github.com/aceld/zinx/ziface"
)

func RegisterHandlers(server ziface.IServer) {
	user.RegisterUserHandlers(server) // register user handlers
}
