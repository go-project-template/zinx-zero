package playerBag

import (
	"sync"
	"zinx-zero/apps/acommon/cfg"
	"zinx-zero/apps/acommon/globalkey"
	"zinx-zero/apps/gamex/internal/ice"
	"zinx-zero/apps/gamex/msg"

	"github.com/zeromicro/go-zero/core/logx"
)

// Check interface implementation.
var _ ice.IPlayerBag = (*PlayerBag)(nil)

func NewPlayerBag(player ice.IPlayer) ice.IPlayerBag {
	return &PlayerBag{}
}

type PlayerBag struct {
	*msg.DBPlayerBag
	sync.RWMutex

	player ice.IPlayer
}

// Init implements ice.IPlayerBag.
func (a *PlayerBag) Init(dbPlayerBag *msg.DBPlayerBag) {
	if dbPlayerBag == nil {
		logx.Errorf("Init PlayerBag dbPlayerBag is nil")
		return
	}
	a.DBPlayerBag = dbPlayerBag
}

// AddItemByItemId implements ice.IPlayerBag.
func (a *PlayerBag) AddItemByItemId(itemId int32, changeCount int64, changeType globalkey.PlayerBag_ItemChangeType) {
	cfg.GetItemByID(itemId)
}

// DelItemByItemId implements ice.IPlayerBag.
func (a *PlayerBag) DelItemByItemId(itemId int32, changeCount int64, changeType globalkey.PlayerBag_ItemChangeType) {
	panic("unimplemented")
}

// DelItemByUniqueId implements ice.IPlayerBag.
func (a *PlayerBag) DelItemByUniqueId(uniqueId int64, changeCount int64, changeType globalkey.PlayerBag_ItemChangeType) {
	panic("unimplemented")
}

// GetAllItem implements ice.IPlayerBag.
func (a *PlayerBag) GetAllItem() (itemInfoList []*msg.ItemInfo) {
	a.doRead(func() {
		itemInfoList = a.GetItemList()
	})
	return itemInfoList
}

// GetItemByUniqueId implements ice.IPlayerBag.
func (a *PlayerBag) GetItemByUniqueId(uniqueId int64) (itemInfo *msg.ItemInfo) {
	a.doRead(func() {
		for _, v := range a.GetItemList() {
			if uniqueId == v.GetUniqueId() {
				itemInfo = v
				break
			}
		}
	})
	return itemInfo
}

// GetItemListByItemId implements ice.IPlayerBag.
func (a *PlayerBag) GetItemListByItemId(itemId int32) (itemInfoList []*msg.ItemInfo) {
	a.doRead(func() {
		for _, v := range a.GetItemList() {
			if itemId == v.GetItemId() {
				itemInfoList = append(itemInfoList, v)
				break
			}
		}
	})
	return itemInfoList
}

func (a *PlayerBag) doWrite(fn func()) {
	a.Lock()
	defer a.Unlock()
	fn()
}

func (a *PlayerBag) doRead(fn func()) {
	a.RLock()
	defer a.RUnlock()
	fn()
}
