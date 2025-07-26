package scenes

import "github.com/hajimehoshi/ebiten/v2"

type SceneId uint

const (
	GameSceneId SceneId = iota
	StartSceneId
	ExitSceneId
	PlayerTestSceneId
	DebugSceneId
	TestLevelSceneId
	SortingSpritesSceneId
)

type Scene interface {
	Update() SceneId
	Draw(screen *ebiten.Image)
	FirstLoad()
	OnEnter()
	OnExit()
	IsLoaded() bool
}
