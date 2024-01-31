package ice

import (
	"zinx-zero/apps/gamex/proto/msg"

	"github.com/aceld/zinx/ziface"
	"google.golang.org/protobuf/proto"
)

type IPlayerManager interface {
	AddPlayer(player IPlayer)
	GetPlayerByRoleId(roleId int64) (player IPlayer, err error)
	GetPlayerByRoleIdStr(roleIdStr string) (player IPlayer, err error)
	RemovePlayer(player IPlayer)
}

type IPlayer interface {
	DoWriteLock(fn func())
	DoReadLock(fn func())
	IPlayerBag
	InitPlayerByDb(dbPlayer *msg.DBPlayer)
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
	SendMsg(msgID msg.MsgId, data proto.Message)
	SendBuffMsg(msgID msg.MsgId, data proto.Message)

	SetIntAttr(k int32, v int64)
	GetIntAttr(k int32) (v int64)
	SetStrAttr(k int32, v string)
	GetStrAttr(k int32) (v string)
}

type IPlayerBag interface {
	AddItemByItemId(itemId int32, changeCount int64, changeType msg.EnumItemChangeType) (code msg.EnumCode)
	DelItemByItemId(itemId int32, changeCount int64, changeType msg.EnumItemChangeType)
	DelItemByUniqueId(uniqueId int64, changeCount int64, changeType msg.EnumItemChangeType)
	GetItemByUniqueId(uniqueId int64) (itemInfo *msg.ItemInfo)
	GetItemByItemId(itemId int32) (itemInfo *msg.ItemInfo)
	GetItemListByItemId(itemId int32) (itemInfoList []*msg.ItemInfo)
	GetAllItem() (itemInfoList []*msg.ItemInfo)
}
