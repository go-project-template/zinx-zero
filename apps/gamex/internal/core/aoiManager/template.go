package aoiManager

import (
	"sync"
	"zinx-zero/apps/gamex/internal/ice"

	"github.com/spf13/cast"
)

// Check interface implementation.
var _ ice.IAoi = (*Aoi)(nil)

type Aoi struct {
	sync.RWMutex

	aoiId    int64
	aoiIdStr string

	MinX  int           // Left boundary coordinate of the area(区域左边界坐标)
	MaxX  int           // Right boundary coordinate of the area(区域右边界坐标)
	CntsX int           // Number of grids in the x direction(x方向格子的数量)
	MinY  int           // Upper boundary coordinate of the area(区域上边界坐标)
	MaxY  int           // Lower boundary coordinate of the area(区域下边界坐标)
	CntsY int           // Number of grids in the y direction(y方向的格子数量)
	grIDs map[int]*GrID // Which grids are present in the current area, key = grid ID, value = grid object(当前区域中都有哪些格子，key=格子ID， value=格子对象)

}

// GetAoiId implements ice.IAoi.
func (a *Aoi) GetAoiId() (aoiId int64) {
	a.doRead(func() {
		aoiId = a.aoiId
	})
	return aoiId
}

// GetAoiIdStr implements ice.IAoi.
func (a *Aoi) GetAoiIdStr() (aoiIdStr string) {
	a.doRead(func() {
		aoiIdStr = a.aoiIdStr
	})
	return aoiIdStr
}

// SetAoiId implements ice.IAoi.
func (a *Aoi) SetAoiId(aoiId int64) {
	a.doWrite(func() {
		a.aoiId = aoiId
		a.aoiIdStr = cast.ToString(aoiId)
	})
}

func (a *Aoi) doWrite(fn func()) {
	a.Lock()
	defer a.Unlock()
	fn()
}

func (a *Aoi) doRead(fn func()) {
	a.RLock()
	defer a.RUnlock()
	fn()
}
