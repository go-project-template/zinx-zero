package sceneManage

import (
	"zinx-zero/apps/gamex/internal/ice"

	"github.com/spf13/cast"
)

// Check interface implementation.
var _ ice.IScene = (*Scene)(nil)

type Scene struct {
	sceneId    int64
	sceneIdStr string
}

// GetSceneId implements ice.IScene.
func (a *Scene) GetSceneId() (sceneId int64) {
	return a.sceneId
}

// GetSceneIdStr implements ice.IScene.
func (a *Scene) GetSceneIdStr() (sceneIdStr string) {
	return a.sceneIdStr
}

// SetSceneId implements ice.IScene.
func (a *Scene) SetSceneId(sceneId int64) {
	a.sceneId = sceneId
	a.sceneIdStr = cast.ToString(sceneId)
}
