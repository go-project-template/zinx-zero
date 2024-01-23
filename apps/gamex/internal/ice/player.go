package ice

import (
	"github.com/aceld/zinx/ziface"
	"google.golang.org/protobuf/proto"
)

type IPlayer interface {
	SetRoleId(roleId int64)
	GetRoleId() (roleId int64)
	GetRoleIdStr() (roleIdStr string)
	SetAccountId(accountId int64)
	GetAccountId() (accountId int64)
	GetAccountIdStr() (accountIdStr string)
	SetNickname(nickname string)
	GetNickname() (nickname string)
	SetConn(conn ziface.IConnection)
	GetConn() (conn ziface.IConnection)
	SendMsg(msgID uint32, data proto.Message)
	SendBuffMsg(msgID uint32, data proto.Message)

	InitPosition()
}

type IPlayerManager interface {
	NewPlayer(roleId int64, conn ziface.IConnection) (player IPlayer)
	AddPlayer(player IPlayer)
	GetPlayerByRoleId(roleId int64) (player IPlayer, err error)
	GetPlayerByRoleIdStr(roleIdStr string) (player IPlayer, err error)
	RemovePlayer(player IPlayer)
}
