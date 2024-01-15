package ice

type IScene interface {
	SetSceneId(sceneId int64)
	GetSceneId() (sceneId int64)
	GetSceneIdStr() (sceneIdStr string)
}

type ISceneManager interface {
	NewScene(id int64) (scene IScene)
	AddScene(scene IScene)
	GetSceneBySceneId(sceneId int64) (scene IScene, err error)
	GetSceneBySceneIdStr(sceneIdStr string) (scene IScene, err error)
	RemoveScene(scene IScene)
}
