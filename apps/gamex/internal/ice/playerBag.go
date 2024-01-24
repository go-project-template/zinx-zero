package ice

import (
	"zinx-zero/apps/acommon/globalkey"
	"zinx-zero/apps/gamex/msg"
)

type IPlayerBag interface {
	Init(dbPlayerBag *msg.DBPlayerBag)
	AddItemByItemId(itemId int64, changeCount int64, changeType globalkey.PlayerBag_ItemChangeType)
	DelItemByItemId(itemId int64, changeCount int64, changeType globalkey.PlayerBag_ItemChangeType)
	DelItemByUniqueId(uniqueId int64, changeCount int64, changeType globalkey.PlayerBag_ItemChangeType)
	GetItemByUniqueId(uniqueId int64) (itemInfo *msg.ItemInfo)
	GetItemListByItemId(itemId int64) (itemInfoList []*msg.ItemInfo)
	GetAllItem() (itemInfoList []*msg.ItemInfo)
}
