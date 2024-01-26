// Code generated by exceltool. DO NOT EDIT!
package cfg

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
)

type Item struct {
	Id        int32  `json:"Id"`        // Item
	Name      string `json:"Name"`      // 名字
	Desc      string `json:"Desc"`      // 描述
	Types     int32  `json:"Types"`     // 类型
	BindType  int32  `json:"BindType"`  // 绑定类型
	StackType int32  `json:"StackType"` // 叠加类型
	StackSize int64  `json:"StackSize"` // 最大叠加数量
	Level     int32  `json:"Level"`     // 等级
	Quality   int32  `json:"Quality"`   // 品质
	EquipSlot int32  `json:"EquipSlot"` // 装备位置
	Icon      string `json:"Icon"`      // 图标
}

var ItemMap map[int32]*Item
var ItemAry []*Item

func initItem() {
	fileName := "./conf/game/Item.json"
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err := json.Unmarshal(bytes, &ItemMap); err != nil {
		panic(err)
	}
	ItemAry = make([]*Item, 0, len(ItemMap))
	for _, item := range ItemMap {
		ItemAry = append(ItemAry, item)
	}
}

func GetItemMap() map[int32]*Item {
	return ItemMap
}

func GetItemAry() []*Item {
	return ItemAry
}

func GetItemByID(id int32) (item *Item) {
	item = ItemMap[id]
	if item == nil {
		logx.Errorf("GetItemByID fail: %d ", id)
	}
	return item
}

func GetItemByIndex(idx int) (item *Item) {
	lens := len(ItemAry)
	if lens <= 0 || idx >= lens {
		logx.Errorf("GetItemByIndex fail: %d ", idx)
		return nil
	}
	return ItemAry[idx]
}