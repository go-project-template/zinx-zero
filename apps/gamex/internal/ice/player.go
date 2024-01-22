package ice

type IPlayer interface {
	SetAccountId(accountId int64)
	GetAccountId() (accountId int64)
	GetAccountIdStr() (userIdStr string)
	SetNickname(nickname string)
	GetNickname() (nickname string)
}

type IPlayerManager interface {
	NewPlayer(accountId int64) (player IPlayer)
	AddPlayer(player IPlayer)
	GetPlayerByAccountId(accountId int64) (player IPlayer, err error)
	GetPlayerByAccountIdStr(userIdStr string) (player IPlayer, err error)
	RemovePlayer(player IPlayer)
}
