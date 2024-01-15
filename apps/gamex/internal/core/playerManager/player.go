package playerManager

import (
	"zinx-zero/apps/gamex/internal/ice"

	"github.com/spf13/cast"
)

// Check interface implementation.
var _ ice.IPlayer = (*Player)(nil)

type Player struct {
	userId    int64
	userIdStr string
	nickname  string
}

// GetNickname implements ice.IPlayer.
func (a *Player) GetNickname() string {
	return a.nickname
}

// GetUserId implements ice.IPlayer.
func (a *Player) GetUserId() int64 {
	return a.userId
}

// GetUserIdStr implements ice.IPlayer.
func (a *Player) GetUserIdStr() (userIdStr string) {
	return a.userIdStr
}

// SetUserId implements ice.IPlayer.
func (a *Player) SetUserId(userId int64) {
	a.userId = userId
	a.userIdStr = cast.ToString(userId)
}

// SetNickname implements ice.IPlayer.
func (a *Player) SetNickname(nickname string) {
	a.nickname = nickname
}
