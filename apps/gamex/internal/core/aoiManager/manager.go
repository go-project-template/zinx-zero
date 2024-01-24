package aoiManager

import (
	"errors"
	"zinx-zero/apps/gamex/internal/ice"

	"github.com/aceld/zinx/zutils"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

// Check interface implementation.
var _ ice.IAoiManager = (*AoiManager)(nil)

var aoiManagerObj = newAoiManager()

func newAoiManager() *AoiManager {
	return &AoiManager{
		aoiMap: zutils.NewShardLockMaps(),
	}
}

func GetAoiManager() ice.IAoiManager {
	return aoiManagerObj
}

type AoiManager struct {
	aoiMap zutils.ShardLockMaps
}

// NewAoi implements ice.IAoiManager.
func (*AoiManager) NewAoi(id int64) (aoi ice.IAoi) {
	aoi = &Aoi{}
	aoi.SetAoiId(id)
	return aoi
}

func (a *AoiManager) AddAoi(aoi ice.IAoi) {
	a.aoiMap.Set(aoi.GetAoiIdStr(), aoi)
	logx.Infof("AddAoi success. %d", aoi.GetAoiId())
}

func (a *AoiManager) GetAoiByAoiId(aoiId int64) (aoi ice.IAoi, err error) {
	return a.GetAoiByAoiIdStr(cast.ToString(aoiId))
}

func (a *AoiManager) GetAoiByAoiIdStr(aoiIdStr string) (aoi ice.IAoi, err error) {
	if conn, ok := a.aoiMap.Get(aoiIdStr); ok {
		return conn.(ice.IAoi), nil
	}
	return nil, errors.New("aoi not found")
}

func (a *AoiManager) RemoveAoi(aoi ice.IAoi) {
	a.aoiMap.Remove(aoi.GetAoiIdStr())
	logx.Infof("RemoveAoi fail. aoiId=%d", aoi.GetAoiId())
}
