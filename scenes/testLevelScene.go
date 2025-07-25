package scenes

import (
	e "game/entities"
	s "game/spritesheets"

	"github.com/hajimehoshi/ebiten/v2"
)

type TestLevelScene struct {
	loaded            bool
	player            *e.Player
	playerSpriteSheet *s.SpriteSheet
	objects           map[string][]e.GameObject
}

func (s *TestLevelScene) GetObjects() *map[string][]e.GameObject {
	return &s.objects
}

func NewTestLevelScene() *TestLevelScene {

	return &TestLevelScene{
		loaded: false,
		player: e.NewPlayer(0, 0),
	}
}

func (d *TestLevelScene) FirstLoad() {
	//read level data
	d.objects = make(map[string][]e.GameObject)
	d.objects["player"] = []e.GameObject{d.player}
	d.objects["enemies"] = []e.GameObject{}
	d.objects["enemyProjectiles"] = []e.GameObject{}
	d.objects["playerProjectiles"] = []e.GameObject{}
	d.objects["staticObjects"] = []e.GameObject{e.NewEntity(e.NewRect(200, 200, 20, 20)), e.NewEntity(e.NewRect(200, 100, 20, 20))}
}

func (d *TestLevelScene) IsLoaded() bool {
	return d.loaded
}

func (d *TestLevelScene) Draw(screen *ebiten.Image) {
	for _, list := range d.objects {
		for _, o := range list {
			o.Draw(screen)
		}
	}

	// ebitenutil.DebugPrint(screen, "Press E to change objects")
}

func (d *TestLevelScene) Update() SceneId {
	d.player.Update(d)
	return TestLevelSceneId
}

func (d *TestLevelScene) OnEnter() {

}

func (d *TestLevelScene) OnExit() {

}

var _ Scene = (*TestLevelScene)(nil)
