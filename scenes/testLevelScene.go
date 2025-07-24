package scenes

import (
	"game/animations"
	"game/entities"
	e "game/entities"
	s "game/spritesheets"
	v "game/utils/math"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	playerSpriteSheet := s.NewSpriteSheet(2, 3, 15, 26)

	return &TestLevelScene{
		loaded: false,
		player: &entities.Player{
			Entity: entities.Entity{
				Sprite: &entities.Sprite{
					Img: playerImg,
					X:   50.0,
					Y:   50.0,
				},
				Collider: &e.Circle{Pos: v.Vec{X: 0, Y: 0}, Radius: 10},
			},
			Animations: map[entities.State]*animations.Animation{
				entities.Up:   animations.NewAnimation(3, 5, 2, 20.0),
				entities.Down: animations.NewAnimation(2, 4, 2, 20.0),
			},
		},
		playerSpriteSheet: playerSpriteSheet,
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
