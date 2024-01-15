package playerManage

import (
	"errors"
	"zinx-zero/apps/gamex/internal/ice"

	"github.com/aceld/zinx/zutils"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

// Check interface implementation.
var _ ice.IPlayerManager = (*PlayerManager)(nil)

var playerManagerObj = newPlayerManager()

func newPlayerManager() *PlayerManager {
	return &PlayerManager{
		playerMap: zutils.NewShardLockMaps(),
	}
}

func GetPlayerManager() ice.IPlayerManager {
	return playerManagerObj
}

type PlayerManager struct {
	playerMap zutils.ShardLockMaps
}

// NewPlayer implements ice.IPlayerManager.
func (*PlayerManager) NewPlayer(userId int64) (player ice.IPlayer) {
	player = &Player{}
	player.SetUserId(userId)
	return player
}

// AddPlayer implements ice.IPlayerManager.
func (a *PlayerManager) AddPlayer(player ice.IPlayer) {
	a.playerMap.Set(player.GetUserIdStr(), player)
	logx.Infof("player add to playerManager successfully: %v", player.GetUserId())
}

// GetPlayerByUserId implements ice.IPlayerManager.
func (a *PlayerManager) GetPlayerByUserId(userId int64) (player ice.IPlayer, err error) {
	return a.GetPlayerByUserIdStr(cast.ToString(userId))
}

// GetPlayerByUserIdStr implements ice.IPlayerManager.
func (a *PlayerManager) GetPlayerByUserIdStr(userIdStr string) (player ice.IPlayer, err error) {
	if conn, ok := a.playerMap.Get(userIdStr); ok {
		return conn.(ice.IPlayer), nil
	}
	return nil, errors.New("player not found")
}

// RemovePlayer implements ice.IPlayerManager.
func (a *PlayerManager) RemovePlayer(player ice.IPlayer) {
	a.playerMap.Remove(player.GetUserIdStr())
	logx.Infof("player Remove userId=%d successfully", player.GetUserId())
}
