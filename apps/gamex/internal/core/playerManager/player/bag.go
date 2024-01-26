package player

import (
	"zinx-zero/apps/acommon/cfg"
	"zinx-zero/apps/gamex/proto/msg"

	"github.com/spf13/cast"
)

func (a *playerImpl) addItemByItemId(itemId int32, changeCount int64,
	changeType msg.EnumItemChangeType) (code msg.EnumCode) {
	code = msg.EnumCode_Code_Fail
	if changeCount <= 0 {
		return
	}
	cfgItem := cfg.GetItemByID(itemId)
	if cfgItem == nil {
		code = msg.EnumCode_Code_Item_NotExist
		return
	}
	var itemInfo *msg.ItemInfo
	switch cfgItem.StackType {
	case cfg.Item_StackType_Stackable:
		itemInfo = a.getItemByItemId(itemId)
	case cfg.Item_StackType_NoStackable:
		for i := 0; i < int(changeCount); i++ {
			itemInfo, code = a.createItemByItemId(itemId)
			if code != msg.EnumCode_Code_Success {
				return
			}
		}

	default:
		code = msg.EnumCode_Code_Item_UnknowStackType
		return
	}
	if itemInfo == nil {
		itemInfo, code = a.createItemByItemId(itemId)
		if code != msg.EnumCode_Code_Success {
			return
		}
	}
	code = msg.EnumCode_Code_Success
	return
}

func (a *playerImpl) delItemByItemId(itemId int32, changeCount int64, changeType msg.ItemChangeType) {
	panic("unimplemented")
}

func (a *playerImpl) delItemByUniqueId(uniqueId int64, changeCount int64, changeType msg.ItemChangeType) {
	panic("unimplemented")
}

func (a *playerImpl) getAllItem() (itemInfoList []*msg.ItemInfo) {
	return a.GetDBPlayerBag().GetItemList()
}

func (a *playerImpl) getItemByUniqueId(uniqueId int64) (itemInfo *msg.ItemInfo) {
	for _, v := range a.GetDBPlayerBag().GetItemList() {
		if uniqueId == v.GetUniqueId() {
			return v
		}
	}
	return itemInfo
}

func (a *playerImpl) getItemByItemId(itemId int32) (itemInfo *msg.ItemInfo) {
	for _, v := range a.GetDBPlayerBag().GetItemList() {
		if itemId == v.GetItemId() {
			itemInfo = v
			break
		}
	}
	return itemInfo
}

func (a *playerImpl) getItemListByItemId(itemId int32) (itemInfoList []*msg.ItemInfo) {
	for _, v := range a.GetDBPlayerBag().GetItemList() {
		if itemId == v.GetItemId() {
			itemInfoList = append(itemInfoList, v)
			break
		}
	}
	return itemInfoList
}

func (a *playerImpl) createItemByItemId(itemId int32) (itemInfo *msg.ItemInfo, code msg.EnumCode) {
	cfgItem := cfg.GetItemByID(itemId)
	if cfgItem == nil {
		code = msg.EnumCode_Code_Item_NotExist
		return
	}
	// 道具唯一id自增
	autoItemId := a.getIntAttr(int32(msg.Enum_PlayerIntAttr_PlayerIntAttr_ItemId))
	autoItemId += 1
	a.setIntAttr(int32(msg.Enum_PlayerIntAttr_PlayerIntAttr_ItemId), autoItemId)
	itemInfo = &msg.ItemInfo{
		UniqueId:  cast.ToInt64(cast.ToString(a.getRoleId()) + cast.ToString(autoItemId)),
		ItemId:    itemId,
		ItemCount: 0,
	}
	code = msg.EnumCode_Code_Success
	return
}
