package ice

type ISceneManager interface {
	AddScene(scene IScene)
	GetSceneBySceneId(sceneId int64) (scene IScene, err error)
	GetSceneBySceneIdStr(sceneIdStr string) (scene IScene, err error)
	RemoveScene(scene IScene)
}

type IScene interface {
	DoWriteLock(fn func())
	DoReadLock(fn func())
	SetSceneId(sceneId int64)
	GetSceneId() (sceneId int64)
	GetSceneIdStr() (sceneIdStr string)
}
