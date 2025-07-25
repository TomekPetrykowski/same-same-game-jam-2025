package scenes

import (
	e "game/entities"
	"slices"

	"game/utils/images"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type TestLevelScene struct {
	loaded  bool
	player  *e.Player
	objects e.ObjectsMap
}

func NewTestLevelScene() *TestLevelScene {

	return &TestLevelScene{
		loaded: false,
		player: e.NewPlayer(0, 0),
	}
}

func (s *TestLevelScene) GetObjects() *e.ObjectsMap {
	return &s.objects
}

func (s *TestLevelScene) AddObject(key e.SceneObjectId, object e.GameObject) {
	s.objects[key] = append(s.objects[key], object)
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
	d.objects = make(e.ObjectsMap)
	d.objects[e.PlayerObjectId] = []e.GameObject{d.player}
	d.objects[e.EnemiesObjectId] = []e.GameObject{
		e.NewBombHead(100, 100),
		e.NewShootyEnemy(140, 140),
	}
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

	for key, list := range d.objects {
		for i := 0; i < len(d.objects[key]); {
			if list[i].IsDeleted() {
				d.objects[key] = slices.Delete(d.objects[key], i, i+1)
			} else {
				i++
			}
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
