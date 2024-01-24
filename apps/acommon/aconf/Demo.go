package aconf
// 自动生成模板Demo

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
    "math/rand"
	"sync"
)
type DemoObj struct {
    ProductionId  int32 `json:"productionId"` // 物品ID
    Amount  int32 `json:"amount"` // 消耗物品数量
}

type Demo struct {
    Id  int32 `json:"id"` // Demo
    Name  string `json:"name"` // 名字
    StrAry  []string `json:"strAry"` // 数组
    Num  int32 `json:"num"` // 数字
    CanSee  int32 `json:"canSee"` // 是否可见
    MoreAry  [][]int32 `json:"moreAry"` // 二维数组
    Obj  map[int32]DemoObj `json:"obj"` // 对象
}

var DemoMap map[int32]Demo
var DemoAry []Demo

func initDemo() {
	v := viper.New()
	v.SetConfigFile("./conf/game/Demo.json")
	v.SetConfigType("json")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(&DemoMap); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&DemoMap); err != nil {
		panic(err)
	}
	DemoAry = make([]Demo, 0, len(DemoMap))
	for _, item := range DemoMap {
        DemoAry = append(DemoAry, item)
    }
}

var loadDemo sync.Once

func GetDemoMap() map[int32]Demo {
	loadDemo.Do(initDemo)
	return DemoMap
}

func GetDemoAry() []Demo {
	loadDemo.Do(initDemo)
	return DemoAry
}

func GetDemoByID(id int32) (Demo, bool) {
	loadDemo.Do(initDemo)
	item,ok := DemoMap[id]
	return item,ok
}

func GetDemoByIndex(idx int) (item Demo,ok bool) {
	loadDemo.Do(initDemo)
	lens := len(DemoAry)
	if lens <=0 || idx >= lens {
	    return
	}
	return DemoAry[idx],true
}

func GetRandDemo() (item Demo,ok bool) {
	loadDemo.Do(initDemo)
	lens := len(DemoAry)
	if lens <=0 {
	    return
	}
	return DemoAry[rand.Intn(lens)],true
}