package player

import (
	"zinx-zero/apps/gamex/proto/msg"

	"github.com/aceld/zinx/ziface"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"

	"github.com/spf13/cast"
)

func NewPlayerImpl(roleId int64, conn ziface.IConnection) *playerImpl {
	player := &playerImpl{}
	player.setRoleId(roleId)
	player.setConn(conn)
	player.initPlayerByDb(&msg.DBPlayer{})
	return player
}

// Logic is handled here to avoid deadlock.
type playerImpl struct {
	*msg.DBPlayer

	roleIdStr    string
	accountIdStr string
	conn         ziface.IConnection
}

func (a *playerImpl) initPlayerByDb(dbPlayer *msg.DBPlayer) {
	if dbPlayer == nil {
		logx.Errorf("Init Player dbPlayer is nil")
		return
	}
	// DBPlayer
	{
		if dbPlayer.IntAttr == nil {
			dbPlayer.IntAttr = make(map[int32]int64)
		}
		if dbPlayer.StrAttr == nil {
			dbPlayer.StrAttr = make(map[int32]string)
		}
	}
	// DBPlayerBag
	{
		if dbPlayer.DBPlayerBag == nil {
			dbPlayer.DBPlayerBag = &msg.DBPlayerBag{}
		}
	}
	a.DBPlayer = dbPlayer
}

// GetIntAttr implements ice.IPlayer.
func (a *playerImpl) getIntAttr(k int32) (v int64) {
	return a.IntAttr[k]
}

// GetStrAttr implements ice.IPlayer.
func (a *playerImpl) getStrAttr(k int32) (v string) {
	return a.StrAttr[k]
}

func (a *playerImpl) setIntAttr(k int32, v int64) {
	a.IntAttr[k] = v
}

func (a *playerImpl) setStrAttr(k int32, v string) {
	a.StrAttr[k] = v
}

func (a *playerImpl) setConn(conn ziface.IConnection) {
	a.conn = conn
}

func (a *playerImpl) getConn() (conn ziface.IConnection) {
	return a.conn
}

func (a *playerImpl) setRoleId(roleId int64) {
	a.RoleId = roleId
	a.roleIdStr = cast.ToString(roleId)
}

func (a *playerImpl) getRoleId() (roleId int64) {
	return a.RoleId
}

func (a *playerImpl) getRoleIdStr() (roleIdStr string) {
	return a.roleIdStr
}

func (a *playerImpl) getNickname() (nickname string) {
	return a.Nickname
}

func (a *playerImpl) getAccountId() (accountId int64) {
	return a.AccountId
}

func (a *playerImpl) getAccountIdStr() (accountIdStr string) {
	return a.accountIdStr
}

func (a *playerImpl) setAccountId(accountId int64) {
	a.AccountId = accountId
	a.accountIdStr = cast.ToString(accountId)
}

func (a *playerImpl) setNickname(nickname string) {
	a.Nickname = nickname
}

func (a *playerImpl) sendMsg(msgID msg.MsgId, data proto.Message, isBuff bool) {
	if a.conn == nil {
		logx.Errorf("SendMsg roleId=%v connection in player is nil", a.getRoleId())
		return
	}
	// Serialize the proto Message structure
	// 将proto Message结构体序列化
	msg, err := proto.Marshal(data)
	if err != nil {
		logx.Errorf("SendMsg roleId=%v marshal msg err: %v", a.getRoleId(), err)
		return
	}
	// Call the Zinx framework's SendMsg to send the packet
	// 调用Zinx框架的SendMsg发包
	if isBuff {
		if err := a.conn.SendBuffMsg(uint32(msgID), msg); err != nil {
			logx.Errorf("SendBuffMsg roleId=%v err: %v", a.getRoleId(), err)
			return
		}
	} else {
		if err := a.conn.SendMsg(uint32(msgID), msg); err != nil {
			logx.Errorf("SendMsg roleId=%v err: %v", a.getRoleId(), err)
			return
		}
	}
}
