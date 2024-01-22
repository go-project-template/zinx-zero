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

// NewPlayer implements ice.IPlayerManager.
func (*PlayerManager) NewPlayer(accountId int64) (player ice.IPlayer) {
	player = &Player{}
	player.SetAccountId(accountId)
	return player
}

// AddPlayer implements ice.IPlayerManager.
func (a *PlayerManager) AddPlayer(player ice.IPlayer) {
	a.playerMap.Set(player.GetAccountIdStr(), player)
	logx.Infof("player add to playerManager successfully: %v", player.GetAccountId())
}

// GetPlayerByAccountId implements ice.IPlayerManager.
func (a *PlayerManager) GetPlayerByAccountId(accountId int64) (player ice.IPlayer, err error) {
	return a.GetPlayerByAccountIdStr(cast.ToString(accountId))
}

// GetPlayerByAccountIdStr implements ice.IPlayerManager.
func (a *PlayerManager) GetPlayerByAccountIdStr(userIdStr string) (player ice.IPlayer, err error) {
	if conn, ok := a.playerMap.Get(userIdStr); ok {
		return conn.(ice.IPlayer), nil
	}
	return nil, errors.New("player not found")
}

// RemovePlayer implements ice.IPlayerManager.
func (a *PlayerManager) RemovePlayer(player ice.IPlayer) {
	a.playerMap.Remove(player.GetAccountIdStr())
	logx.Infof("player Remove accountId=%d successfully", player.GetAccountId())
}
