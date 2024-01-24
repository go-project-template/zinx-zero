package playerManager

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

// AddPlayer implements ice.IPlayerManager.
func (a *PlayerManager) AddPlayer(player ice.IPlayer) {
	a.playerMap.Set(player.GetRoleIdStr(), player)
	logx.Infof("player add to playerManager successfully: %v", player.GetRoleId())
}

// GetPlayerByRoleId implements ice.IPlayerManager.
func (a *PlayerManager) GetPlayerByRoleId(roleId int64) (player ice.IPlayer, err error) {
	return a.GetPlayerByRoleIdStr(cast.ToString(roleId))
}

// GetPlayerByRoleIdStr implements ice.IPlayerManager.
func (a *PlayerManager) GetPlayerByRoleIdStr(roleIdStr string) (player ice.IPlayer, err error) {
	if conn, ok := a.playerMap.Get(roleIdStr); ok {
		return conn.(ice.IPlayer), nil
	}
	return nil, errors.New("player not found")
}

// RemovePlayer implements ice.IPlayerManager.
func (a *PlayerManager) RemovePlayer(player ice.IPlayer) {
	a.playerMap.Remove(player.GetRoleIdStr())
	logx.Infof("player Remove roleId=%d successfully", player.GetRoleId())
}
