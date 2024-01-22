package playerManager

import (
	"zinx-zero/apps/gamex/internal/ice"

	"github.com/spf13/cast"
)

// Check interface implementation.
var _ ice.IPlayer = (*Player)(nil)

type Player struct {
	accountId int64
	userIdStr string
	nickname  string
}

// GetNickname implements ice.IPlayer.
func (a *Player) GetNickname() string {
	return a.nickname
}

// GetAccountId implements ice.IPlayer.
func (a *Player) GetAccountId() int64 {
	return a.accountId
}

// GetAccountIdStr implements ice.IPlayer.
func (a *Player) GetAccountIdStr() (userIdStr string) {
	return a.userIdStr
}

// SetAccountId implements ice.IPlayer.
func (a *Player) SetAccountId(accountId int64) {
	a.accountId = accountId
	a.userIdStr = cast.ToString(accountId)
}

// SetNickname implements ice.IPlayer.
func (a *Player) SetNickname(nickname string) {
	a.nickname = nickname
}
