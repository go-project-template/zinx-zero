package sceneManager

import (
	"sync"
	"zinx-zero/apps/gamex/internal/ice"

	"github.com/spf13/cast"
)

// Check interface implementation.
var _ ice.IScene = (*Scene)(nil)

func NewScene(id int64) ice.IScene {
	scene := &Scene{}
	scene.SetSceneId(id)
	return scene
}

type Scene struct {
	sync.RWMutex

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

func (a *Scene) DoWriteLock(fn func()) {
	a.Lock()
	defer a.Unlock()
	fn()
}

func (a *Scene) DoReadLock(fn func()) {
	a.RLock()
	defer a.RUnlock()
	fn()
}
