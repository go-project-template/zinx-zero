package sceneManager

import (
	"errors"
	"zinx-zero/apps/gamex/internal/ice"

	"github.com/aceld/zinx/zutils"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

// Check interface implementation.
var _ ice.ISceneManager = (*SceneManager)(nil)

var sceneManagerObj = newSceneManager()

func newSceneManager() *SceneManager {
	return &SceneManager{
		sceneMap: zutils.NewShardLockMaps(),
	}
}

func GetSceneManager() ice.ISceneManager {
	return sceneManagerObj
}

type SceneManager struct {
	sceneMap zutils.ShardLockMaps
}

// NewScene implements ice.ISceneManager.
func (*SceneManager) NewScene(id int64) (scene ice.IScene) {
	scene = &Scene{}
	scene.SetSceneId(id)
	return scene
}

func (a *SceneManager) AddScene(scene ice.IScene) {
	a.sceneMap.Set(scene.GetSceneIdStr(), scene)
	logx.Infof("scene add to sceneManager successfully: %v", scene.GetSceneId())
}

func (a *SceneManager) GetSceneBySceneId(sceneId int64) (scene ice.IScene, err error) {
	return a.GetSceneBySceneIdStr(cast.ToString(sceneId))
}

func (a *SceneManager) GetSceneBySceneIdStr(sceneIdStr string) (scene ice.IScene, err error) {
	if conn, ok := a.sceneMap.Get(sceneIdStr); ok {
		return conn.(ice.IScene), nil
	}
	return nil, errors.New("scene not found")
}

func (a *SceneManager) RemoveScene(scene ice.IScene) {
	a.sceneMap.Remove(scene.GetSceneIdStr())
	logx.Infof("scene Remove sceneId=%d successfully", scene.GetSceneId())
}
