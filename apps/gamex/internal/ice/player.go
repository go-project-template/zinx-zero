package ice

import (
	"zinx-zero/apps/gamex/msg"
	"zinx-zero/apps/gamex/pb"

	"github.com/aceld/zinx/ziface"
	"google.golang.org/protobuf/proto"
)

type IPlayer interface {
	Init(dbPlayer *msg.DBPlayer)
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
	SendMsg(msgID pb.MsgId, data proto.Message)
	SendBuffMsg(msgID pb.MsgId, data proto.Message)
}

type IPlayerManager interface {
	AddPlayer(player IPlayer)
	GetPlayerByRoleId(roleId int64) (player IPlayer, err error)
	GetPlayerByRoleIdStr(roleIdStr string) (player IPlayer, err error)
	RemovePlayer(player IPlayer)
}
