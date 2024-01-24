package sceneManager

import (
	"sync"
	"zinx-zero/apps/gamex/internal/ice"

	"github.com/spf13/cast"
)

// Check interface implementation.
var _ ice.IScene = (*Scene)(nil)

func NewScene(id int64) (scene ice.IScene) {
	scene = &Scene{}
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
	a.doRead(func() {
		sceneId = a.sceneId
	})
	return sceneId
}

// GetSceneIdStr implements ice.IScene.
func (a *Scene) GetSceneIdStr() (sceneIdStr string) {
	a.doRead(func() {
		sceneIdStr = a.sceneIdStr
	})
	return sceneIdStr
}

// SetSceneId implements ice.IScene.
func (a *Scene) SetSceneId(sceneId int64) {
	a.doWrite(func() {
		a.sceneId = sceneId
		a.sceneIdStr = cast.ToString(sceneId)
	})
}

func (a *Scene) doWrite(fn func()) {
	a.Lock()
	defer a.Unlock()
	fn()
}

func (a *Scene) doRead(fn func()) {
	a.RLock()
	defer a.RUnlock()
	fn()
}
