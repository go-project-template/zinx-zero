package player

import (
	"sync"
	"zinx-zero/apps/gamex/internal/ice"
	"zinx-zero/apps/gamex/proto/msg"

	"github.com/aceld/zinx/ziface"
	"google.golang.org/protobuf/proto"
)

// Check interface implementation.
var _ ice.IPlayer = (*Player)(nil)

func NewPlayer(roleId int64, conn ziface.IConnection) ice.IPlayer {
	player := &Player{
		playerImpl: NewPlayerImpl(roleId, conn),
	}
	return player
}

type Player struct {
	sync.RWMutex

	playerImpl *playerImpl
}

// GetIntAttr implements ice.IPlayer.
func (a *Player) GetIntAttr(k int32) (v int64) {
	a.DoReadLock(func() { v = a.playerImpl.getIntAttr(k) })
	return v
}

// GetStrAttr implements ice.IPlayer.
func (a *Player) GetStrAttr(k int32) (v string) {
	a.DoReadLock(func() { v = a.playerImpl.getStrAttr(k) })
	return v
}

// SetIntAttr implements ice.IPlayer.
func (a *Player) SetIntAttr(k int32, v int64) {
	a.DoWriteLock(func() { a.playerImpl.setIntAttr(k, v) })
}

// SetStrAttr implements ice.IPlayer.
func (a *Player) SetStrAttr(k int32, v string) {
	a.DoWriteLock(func() { a.playerImpl.setStrAttr(k, v) })
}

// GetAccountId implements ice.IPlayer.
func (a *Player) GetAccountId() (accountId int64) {
	a.DoReadLock(func() { accountId = a.playerImpl.getAccountId() })
	return accountId
}

// GetNickname implements ice.IPlayer.
func (a *Player) GetNickname() (nickname string) {
	a.DoReadLock(func() { nickname = a.playerImpl.getNickname() })
	return nickname
}

// GetRoleId implements ice.IPlayer.
func (a *Player) GetRoleId() (roleId int64) {
	a.DoReadLock(func() { roleId = a.playerImpl.getRoleId() })
	return roleId
}

// AddItemByItemId implements ice.IPlayer.
func (a *Player) AddItemByItemId(itemId int32, changeCount int64, changeType msg.EnumItemChangeType) (code msg.EnumCode) {
	a.DoWriteLock(func() { code = a.playerImpl.addItemByItemId(itemId, changeCount, changeType) })
	return code
}

// DelItemByItemId implements ice.IPlayer.
func (a *Player) DelItemByItemId(itemId int32, changeCount int64, changeType msg.EnumItemChangeType) {
	panic("unimplemented")
}

// DelItemByUniqueId implements ice.IPlayer.
func (a *Player) DelItemByUniqueId(uniqueId int64, changeCount int64, changeType msg.EnumItemChangeType) {
	panic("unimplemented")
}

// GetAccountIdStr implements ice.IPlayer.
func (a *Player) GetAccountIdStr() (accountIdStr string) {
	a.DoReadLock(func() { accountIdStr = a.playerImpl.getAccountIdStr() })
	return accountIdStr
}

// GetAllItem implements ice.IPlayer.
func (a *Player) GetAllItem() (itemInfoList []*msg.ItemInfo) {
	a.DoReadLock(func() { itemInfoList = a.playerImpl.getAllItem() })
	return itemInfoList
}

// GetConn implements ice.IPlayer.
func (a *Player) GetConn() (conn ziface.IConnection) {
	a.DoReadLock(func() { conn = a.playerImpl.getConn() })
	return conn
}

// GetItemByItemId implements ice.IPlayer.
func (a *Player) GetItemByItemId(itemId int32) (itemInfo *msg.ItemInfo) {
	a.DoReadLock(func() { itemInfo = a.playerImpl.getItemByItemId(itemId) })
	return itemInfo
}

// GetItemByUniqueId implements ice.IPlayer.
func (a *Player) GetItemByUniqueId(uniqueId int64) (itemInfo *msg.ItemInfo) {
	a.DoReadLock(func() { itemInfo = a.playerImpl.getItemByUniqueId(uniqueId) })
	return itemInfo
}

// GetItemListByItemId implements ice.IPlayer.
func (a *Player) GetItemListByItemId(itemId int32) (itemInfoList []*msg.ItemInfo) {
	a.DoReadLock(func() { itemInfoList = a.playerImpl.getItemListByItemId(itemId) })
	return itemInfoList
}

// GetRoleIdStr implements ice.IPlayer.
func (a *Player) GetRoleIdStr() (roleIdStr string) {
	a.DoReadLock(func() { roleIdStr = a.playerImpl.getRoleIdStr() })
	return roleIdStr
}

// InitPlayerByDb implements ice.IPlayer.
func (a *Player) InitPlayerByDb(dbPlayer *msg.DBPlayer) {
	a.DoWriteLock(func() { a.playerImpl.initPlayerByDb(dbPlayer) })
}

// SetAccountId implements ice.IPlayer.
func (a *Player) SetAccountId(accountId int64) {
	a.DoWriteLock(func() { a.playerImpl.setAccountId(accountId) })
}

// SetConn implements ice.IPlayer.
func (a *Player) SetConn(conn ziface.IConnection) {
	a.DoWriteLock(func() { a.playerImpl.setConn(conn) })
}

// SetNickname implements ice.IPlayer.
func (a *Player) SetNickname(nickname string) {
	a.DoWriteLock(func() { a.playerImpl.setNickname(nickname) })
}

// SetRoleId implements ice.IPlayer.
func (a *Player) SetRoleId(roleId int64) {
	a.DoWriteLock(func() { a.playerImpl.setRoleId(roleId) })
}

func (a *Player) SendMsg(msgID msg.MsgId, data proto.Message) {
	a.playerImpl.sendMsg(msgID, data, false)
}
func (a *Player) SendBuffMsg(msgID msg.MsgId, data proto.Message) {
	a.playerImpl.sendMsg(msgID, data, true)
}

func (a *Player) DoWriteLock(fn func()) {
	a.Lock()
	defer a.Unlock()
	fn()
}

func (a *Player) DoReadLock(fn func()) {
	a.RLock()
	defer a.RUnlock()
	fn()
}
