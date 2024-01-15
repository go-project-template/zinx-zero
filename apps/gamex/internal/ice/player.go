package ice

type IPlayer interface {
	SetUserId(userId int64)
	GetUserId() (userId int64)
	GetUserIdStr() (userIdStr string)
	SetNickname(nickname string)
	GetNickname() (nickname string)
}

type IPlayerManager interface {
	NewPlayer(userId int64) (player IPlayer)
	AddPlayer(player IPlayer)
	GetPlayerByUserId(userId int64) (player IPlayer, err error)
	GetPlayerByUserIdStr(userIdStr string) (player IPlayer, err error)
	RemovePlayer(player IPlayer)
}
