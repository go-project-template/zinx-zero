package playerManager

import (
	"sync"
	"zinx-zero/apps/acommon/arand"
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
	X            float32 // Planar x coordinate(平面x坐标)
	Y            float32 // Height(高度)
	Z            float32 // Planar y coordinate (Note: not Y)- 平面y坐标 (注意不是Y)
	V            float32 //  Rotation 0-360 degrees(旋转0-360度)
}

// SyncSurrounding implements ice.IPlayer.
func (a *Player) SyncSurrounding() {
	//1 Get pIDs of players in the surrounding nine grids based on the player's position
	// 根据自己的位置，获取周围九宫格内的玩家pID
	pIDs := WorldMgrObj.AoiMgr.GetPIDsByPos(p.X, p.Z)

	// 2 Get all player objects based on the pIDs
	// 根据pID得到所有玩家对象
	players := make([]*Player, 0, len(pIDs))

	// 3 Send MsgID:200 message to these players to display themselves in each other's views
	// 给这些玩家发送MsgID:200消息，让自己出现在对方视野中
	for _, pID := range pIDs {
		players = append(players, WorldMgrObj.GetPlayerByPID(int32(pID)))
	}

	// 3.1 Assemble MsgID200 proto data
	// 组建MsgID200 proto数据
	x, y, z, v := a.GetPosition()
	msg := &pb.BroadCast{
		PID: int32(a.GetRoleId()),
		Tp:  2, //TP:2 represents broadcasting coordinates (广播坐标)
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: x,
				Y: y,
				Z: z,
				V: v,
			},
		},
	}

	// 3.2 Send the 200 message to each player's client to display characters
	// 每个玩家分别给对应的客户端发送200消息，显示人物
	for _, player := range players {
		player.SendBuffMsg(200, msg)
	}
	// 4 Make surrounding players in the nine grids appear in the player's view
	// 让周围九宫格内的玩家出现在自己的视野中

	// 4.1 Create Message SyncPlayers data
	// 制作Message SyncPlayers 数据
	playersData := make([]*pb.Player, 0, len(players))
	for _, player := range players {
		p := &pb.Player{
			PID: int32(player.GetRoleId()),
			P: &pb.Position{
				X: player.X,
				Y: player.Y,
				Z: player.Z,
				V: player.V,
			},
		}
		playersData = append(playersData, p)
	}

	// 4.2 Encapsulate SyncPlayers protobuf data
	// 封装SyncPlayer protobuf数据
	SyncPlayersMsg := &pb.SyncPlayers{
		Ps: playersData[:],
	}

	// 4.3 Send all player data to the current player to display surrounding players
	// 给当前玩家发送需要显示周围的全部玩家数据
	a.SendBuffMsg(202, SyncPlayersMsg)
}

// GetPosition implements ice.IPlayer.
func (a *Player) GetPosition() (X float32, Y float32, Z float32, V float32) {
	a.doRead(func() {
		X = a.X
		Y = a.Y
		Z = a.Z
		V = a.V
	})
	return
}

// BroadCastStartPosition implements ice.IPlayer.
func (a *Player) BroadCastStartPosition() {
	x, y, z, v := a.GetPosition()
	// Assemble MsgID200 proto data
	// (组建MsgID200 proto数据)
	msg := &pb.BroadCast{
		PID: int32(a.GetRoleId()),
		Tp:  2, //TP:2  represents broadcasting coordinates (广播坐标)
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: x,
				Y: y,
				Z: z,
				V: v,
			},
		},
	}

	// Send data to the client
	// 发送数据给客户端
	a.SendBuffMsg(200, msg)
}

// SyncPID implements ice.IPlayer.
func (a *Player) SyncPID() {
	msg := &pb.SyncPID{PID: int32(a.GetRoleId())}
	a.SendBuffMsg(pb.MsgId_SyncPID_ID, msg)
}

// InitPosition implements ice.IPlayer.
func (a *Player) InitPosition() {
	a.doWrite(func() {
		a.X = float32(160 + arand.Intn(50)) // Randomly offset on the X-axis based on the point 160(随机在160坐标点 基于X轴偏移若干坐标)
		a.Y = 0                             // Height is 0
		a.Z = float32(134 + arand.Intn(50)) // Randomly offset on the Y-axis based on the point 134(随机在134坐标点 基于Y轴偏移若干坐标)
		a.V = 0                             // Angle is 0, not yet implemented(角度为0，尚未实现)
	})
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
