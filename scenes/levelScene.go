package scenes

import (
	e "game/entities"
	"image/color"
	"slices"

	"game/utils/images"
	"game/utils/scene"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type LevelScene struct {
	loaded     bool
	background *ebiten.Image
	objects    e.ObjectsMap
	// waves      []e.Wave
	// delay in frames
	// wavesDelay int
}

func NewLevelScene() *LevelScene {

	return &LevelScene{
		loaded:     false,
		// waves:      []e.Wave{},
		// wavesDelay: 60 * 5,
	}
}

func (s *LevelScene) GetObjects() *e.ObjectsMap {
	return &s.objects
}

func (s *LevelScene) AddObject(key e.SceneObjectId, object e.GameObject) {
	s.objects[key] = append(s.objects[key], object)
}

func (d *LevelScene) FirstLoad() {
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

	

	background := images.LoadImage(
		"assets/background.png",
		"Error while loading player image.",
		&images.DefaultPlaceholder,
	)
	d.background = background

	//read level data
	d.objects = make(map[e.SceneObjectId][]e.GameObject)
	d.objects[e.PortalsObjectId] = []e.GameObject{
		e.NewEntity(e.NewRect(320-10, 39, 20, 20), &placeholderSprite),
		e.NewEntity(e.NewRect(50, 360/2-10, 20, 20), &placeholderSprite),
		e.NewEntity(e.NewRect(640-50-20, 360/2-10, 20, 20), &placeholderSprite),
		e.NewEntity(e.NewRect(320-10, 360-39-20, 20, 20), &placeholderSprite),
	}
	
	d.objects[e.PlayerObjectId] = []e.GameObject{e.NewPlayer(320, 180)}
	d.objects[e.EnemiesObjectId] = []e.GameObject{
		e.NewBombHead(100, 100),
		e.NewShootyEnemy(140, 140),
	}
	d.objects[e.EnemyProjectilesObjectId] = []e.GameObject{}
	d.objects[e.PlayerProjectilesObjectId] = []e.GameObject{}
	d.objects[e.StaticsObjectId] = []e.GameObject{
		e.NewEntity(e.NewRect(27, 4, 586, 32), &e.Sprite{Img: ebiten.NewImage(1, 1)}),
		e.NewEntity(e.NewRect(27, 4, 20, 352), &e.Sprite{Img: ebiten.NewImage(1, 1)}),
		e.NewEntity(e.NewRect(27, 4+352-32-2, 586, 32), &e.Sprite{Img: ebiten.NewImage(1, 1)}),
		e.NewEntity(e.NewRect(27+586-20, 4, 20, 352), &e.Sprite{Img: ebiten.NewImage(1, 1)}),
	}

	// waves
	// d.waves = []e.Wave{
	// 	{
	// 		EnemyTypeQuantity: map[e.EnemyTypeId]int{
	// 			e.Shooter:        2,
	// 			e.BombheadTypeId: 5,
	// 		},
	// 		SpawnDelay: 60,
	// 	},
	// 	{
	// 		EnemyTypeQuantity: map[e.EnemyTypeId]int{
	// 			e.Shooter:        1,
	// 			e.BombheadTypeId: 10,
	// 		},
	// 		SpawnDelay: 30,
	// 	},
	// }

	// d.wavesDelay = 60 * 10
}

func (d *LevelScene) IsLoaded() bool {
	return d.loaded
}

func (s *LevelScene) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(16, 4)
	screen.DrawImage(s.background, op)

	allObjects := scene.GetSortedObjects(&s.objects)

	for _, obj := range *allObjects {
		obj.Draw(screen)
	}

	// ebitenutil.DebugPrint(screen, "Press E to change objects")
}

func (d *LevelScene) Update() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		return PauseSceneId
	}

	// d.wavesDelay -= 1

	// for _, wave := range d.waves {
	// 	if d.wavesDelay <= 0 {
	// 		wave.SpawnEnemies(d)
	// 		d.wavesDelay = 60 * 20
	// 	}
	// }

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
	return LevelSceneId
}

func (d *LevelScene) OnEnter() {

}

func (d *LevelScene) OnExit() {

}

var _ Scene = (*LevelScene)(nil)
