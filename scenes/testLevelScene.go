package scenes

import (
	e "game/entities"

	spritesheet "game/spritesheets"
	"game/utils/images"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type TestLevelScene struct {

	loaded            bool
	player            *e.Player
	objects           map[e.SceneObject][]e.GameObject
}

func (s *TestLevelScene) GetObjects() *map[e.SceneObject][]e.GameObject {
	return &s.objects
}

func (s *TestLevelScene) AddObject(key string, object e.GameObject) {
	s.objects[key] = append(s.objects[key], object)
}

func NewTestLevelScene() *TestLevelScene {

	return &TestLevelScene{
		loaded: false,
		player: e.NewPlayer(0, 0),
	}
}

func (d *TestLevelScene) FirstLoad() {
	// Some default object sprite as placeholder
	placeholderSprite := e.Sprite{
		Img: images.CreatePlaceholderImage(
			&images.PlaceholderImage{
				Width:  20,
				Height: 20,
				Color:  color.RGBA{0, 220, 0, 255},
			},
		),
	}

	//read level data
	d.objects = make(map[e.SceneObject][]e.GameObject)
	d.objects[e.PlayerObjectId] = []e.GameObject{d.player}
	d.objects[e.EnemiesObjectId] = []e.GameObject{e.NewBasicEnemy(200, 300), e.NewShootyEnemy(100, 100)}
	d.objects[e.EnemyProjectilesObjectId] = []e.GameObject{}
	d.objects[e.PlayerProjectilesObjectId] = []e.GameObject{}
	d.objects[e.StaticsObjectId] = []e.GameObject{
		e.NewEntity(
			e.NewRect(200, 200, 20, 20),
			&placeholderSprite,
		),
		e.NewEntity(
			e.NewRect(200, 100, 20, 20),
			&placeholderSprite,
		),
	}
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
	for _, list := range d.objects {
		for _, o := range list {
			o.Update(d)
		}
	}
	// d.player.Update(d)
	return TestLevelSceneId
}

func (d *TestLevelScene) OnEnter() {

}

func (d *TestLevelScene) OnExit() {

}

var _ Scene = (*TestLevelScene)(nil)
