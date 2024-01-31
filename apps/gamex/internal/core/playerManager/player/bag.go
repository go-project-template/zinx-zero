package player

import (
	"github.com/zeromicro/go-zero/core/logx"
	"zinx-zero/apps/acommon/acfg"
	"zinx-zero/apps/acommon/cfg"
	"zinx-zero/apps/gamex/proto/msg"

	"github.com/spf13/cast"
)

func (a *playerImpl) addItemByItemInfo(itemInfo *msg.ItemInfo) {
	a.DBPlayerBag.ItemList = append(a.DBPlayerBag.ItemList, itemInfo)
}

// 添加可叠加的道具
func (a *playerImpl) addStackableItem(cfgItem *cfg.Cfg_Item, changeCount int64) (code msg.EnumCode) {
	code = msg.EnumCode_Code_Fail
	itemId := cfgItem.Id
	var itemInfo = a.getItemByItemId(itemId)
	//创建新的道具
	if itemInfo == nil {
		itemInfo, code = a.createItemByItemId(itemId)
		if code != msg.EnumCode_Code_Success {
			return
		}
		a.addItemByItemInfo(itemInfo)
	}
	// 检查是否达到叠加上限
	if itemInfo.GetCount()+changeCount > cfgItem.StackSize {
		code = msg.EnumCode_Code_Item_ExceedStackSize
		return
	}
	itemInfo.Count += changeCount
	code = msg.EnumCode_Code_Success
	logx.Debugf("添加道具成功", logx.Field("info", itemInfo.String()))
	return
}

func (a *playerImpl) addItemByItemId(itemId int32, changeCount int64,
	changeType msg.EnumItemChangeType) (code msg.EnumCode) {
	code = msg.EnumCode_Code_Fail
	if changeCount <= 0 {
		return
	}
	cfgItem := cfg.GetCfg_ItemById(itemId)
	if cfgItem == nil {
		code = msg.EnumCode_Code_Item_NotExist
		return
	}
	switch cfgItem.StackType {
	case acfg.Item_StackType_Stackable:
		code = a.addStackableItem(cfgItem, changeCount)
	case acfg.Item_StackType_NoStackable:
		for i := 0; i < int(changeCount); i++ {
			itemInfo, code := a.createItemByItemId(itemId)
			if code != msg.EnumCode_Code_Success {
				return code
			}
			a.addItemByItemInfo(itemInfo)
			logx.Debugf("添加道具成功", logx.Field("info", itemInfo.String()))
		}
		code = msg.EnumCode_Code_Success
		return
	default:
		code = msg.EnumCode_Code_Item_UnknowStackType
		return
	}
	return
}

func (a *playerImpl) delItemByItemId(itemId int32, changeCount int64, changeType msg.EnumItemChangeType) {
	panic("unimplemented")
}

func (a *playerImpl) delItemByUniqueId(uniqueId int64, changeCount int64, changeType msg.EnumItemChangeType) {
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
		if itemId == v.GetCfgItemId() {
			itemInfo = v
			break
		}
	}
	return itemInfo
}

func (a *playerImpl) getItemListByItemId(itemId int32) (itemInfoList []*msg.ItemInfo) {
	for _, v := range a.GetDBPlayerBag().GetItemList() {
		if itemId == v.GetCfgItemId() {
			itemInfoList = append(itemInfoList, v)
			break
		}
	}
	return itemInfoList
}

func (a *playerImpl) createItemByItemId(itemId int32) (itemInfo *msg.ItemInfo, code msg.EnumCode) {
	cfgItem := cfg.GetCfg_ItemById(itemId)
	if cfgItem == nil {
		code = msg.EnumCode_Code_Item_NotExist
		return
	}
	// 道具唯一id自增
	autoItemId := a.getIntAttr(int32(msg.EnumPlayerIntAttr_PlayerIntAttr_ItemId))
	autoItemId += 1
	a.setIntAttr(int32(msg.EnumPlayerIntAttr_PlayerIntAttr_ItemId), autoItemId)
	itemInfo = &msg.ItemInfo{
		UniqueId:  cast.ToInt64(cast.ToString(a.getRoleId()) + cast.ToString(autoItemId)),
		CfgItemId: itemId,
		Count:     0,
	}
	code = msg.EnumCode_Code_Success
	return
}
