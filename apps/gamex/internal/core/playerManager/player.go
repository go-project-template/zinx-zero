package playerManager

import (
	"sync"
	"zinx-zero/apps/gamex/internal/ice"
	"zinx-zero/apps/gamex/pb"

	"github.com/aceld/zinx/ziface"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"

	"github.com/spf13/cast"
)

// Check interface implementation.
var _ ice.IPlayer = (*Player)(nil)

type Player struct {
	sync.RWMutex

	roleId       int64
	roleIdStr    string
	accountId    int64
	accountIdStr string
	nickname     string
	conn         ziface.IConnection
}

// SendMsg Send messages to the client, mainly serializing and sending the protobuf data of the pb Message
//
//	(发送消息给客户端，主要是将pb的protobuf数据序列化之后发送)
func (a *Player) SendMsg(msgID pb.MsgId, data proto.Message) {
	if a.conn == nil {
		logx.Errorf("SendMsg roleId=%v connection in player is nil", a.GetRoleId())
		return
	}
	// Serialize the proto Message structure
	// 将proto Message结构体序列化
	msg, err := proto.Marshal(data)
	if err != nil {
		logx.Errorf("SendMsg roleId=%v marshal msg err: %v", a.GetRoleId(), err)
		return
	}

	// Call the Zinx framework's SendMsg to send the packet
	// 调用Zinx框架的SendMsg发包
	if err := a.conn.SendMsg(uint32(msgID), msg); err != nil {
		logx.Errorf("SendMsg roleId=%v err: %v", a.GetRoleId(), err)
		return
	}
}

// SendBuffMsg Send messages to the client, mainly serializing and sending the protobuf data of the pb Message
//
//	(发送消息给客户端，主要是将pb的protobuf数据序列化之后发送)
func (a *Player) SendBuffMsg(msgID pb.MsgId, data proto.Message) {
	if a.conn == nil {
		logx.Errorf("SendBuffMsg roleId=%v connection in player is nil", a.GetRoleId())
		return
	}
	// Serialize the proto Message structure
	// 将proto Message结构体序列化
	msg, err := proto.Marshal(data)
	if err != nil {
		logx.Errorf("SendBuffMsg roleId=%v marshal msg err: %v", a.GetRoleId(), err)
		return
	}

	// Call the Zinx framework's SendMsg to send the packet
	// 调用Zinx框架的SendMsg发包
	if err := a.conn.SendBuffMsg(uint32(msgID), msg); err != nil {
		logx.Errorf("SendBuffMsg roleId=%v err: %v", a.GetRoleId(), err)
		return
	}
}

func (a *Player) SetConn(conn ziface.IConnection) {
	a.doWrite(func() {
		a.conn = conn
	})
}

func (a *Player) GetConn() (conn ziface.IConnection) {
	a.doRead(func() {
		conn = a.conn
	})
	return conn
}

func (a *Player) SetRoleId(roleId int64) {
	a.doWrite(func() {
		a.roleId = roleId
		a.roleIdStr = cast.ToString(roleId)
	})
}

func (a *Player) GetRoleId() (roleId int64) {
	a.doRead(func() {
		roleId = a.roleId
	})
	return roleId
}

func (a *Player) GetRoleIdStr() (roleIdStr string) {
	a.doRead(func() {
		roleIdStr = a.roleIdStr
	})
	return roleIdStr
}

// GetNickname implements ice.IPlayer.
func (a *Player) GetNickname() (nickname string) {
	a.doRead(func() {
		nickname = a.nickname
	})
	return nickname
}

// GetAccountId implements ice.IPlayer.
func (a *Player) GetAccountId() (accountId int64) {
	a.doRead(func() {
		accountId = a.accountId
	})
	return accountId
}

// GetAccountIdStr implements ice.IPlayer.
func (a *Player) GetAccountIdStr() (accountIdStr string) {
	a.doRead(func() {
		accountIdStr = a.accountIdStr
	})
	return accountIdStr
}

// SetAccountId implements ice.IPlayer.
func (a *Player) SetAccountId(accountId int64) {
	a.doWrite(func() {
		a.accountId = accountId
		a.accountIdStr = cast.ToString(accountId)
	})
}

// SetNickname implements ice.IPlayer.
func (a *Player) SetNickname(nickname string) {
	a.doWrite(func() {
		a.nickname = nickname
	})
}

func (a *Player) doWrite(fn func()) {
	a.Lock()
	defer a.Unlock()
	fn()
}

func (a *Player) doRead(fn func()) {
	a.RLock()
	defer a.RUnlock()
	fn()
}
