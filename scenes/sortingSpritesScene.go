package scenes

import (
	e "game/entities"
	"game/utils/images"
	"game/utils/scene"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type SortingSpritesScene struct {
	loaded  bool
	player  *e.Player
	objects e.ObjectsMap
}

func NewSortingSpritesScene() *SortingSpritesScene {
	return &SortingSpritesScene{
		loaded: false,
		player: e.NewPlayer(50, 50),
	}
}

func (s *SortingSpritesScene) FirstLoad() {
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
	s.objects = make(map[e.SceneObjectId][]e.GameObject)
	s.objects[e.PlayerObjectId] = []e.GameObject{s.player}
	s.objects[e.EnemiesObjectId] = []e.GameObject{
		e.NewBasicEnemy(200, 300, 0.7),
		e.NewBasicEnemy(200, 300, 1),
		e.NewShootyEnemy(100, 100),
	}
	s.objects[e.EnemyProjectilesObjectId] = []e.GameObject{}
	s.objects[e.PlayerProjectilesObjectId] = []e.GameObject{}
	s.objects[e.StaticsObjectId] = []e.GameObject{
		e.NewEntity(
			e.NewRect(200, 200, 20, 20),
			&placeholderSprite,
		),
		e.NewEntity(
			e.NewRect(200, 100, 20, 20),
			&placeholderSprite,
		),
	}

	s.loaded = true
}

func (s *SortingSpritesScene) Draw(screen *ebiten.Image) {
	allObjects := scene.GetSortedObjects(&s.objects)

	for _, obj := range *allObjects {
		obj.Draw(screen)
	}
}

func (s *SortingSpritesScene) IsLoaded() bool {
	return s.loaded
}

func (s *SortingSpritesScene) OnEnter() {

}

func (s *SortingSpritesScene) OnExit() {

}

func (s *SortingSpritesScene) GetObjects() *e.ObjectsMap {
	return &s.objects
}

func (s *SortingSpritesScene) AddObject(key e.SceneObjectId, object e.GameObject) {
	s.objects[key] = append(s.objects[key], object)
}

func (s *SortingSpritesScene) Update() SceneId {

	for _, list := range s.objects {
		for _, o := range list {
			o.Update(s)
		}
	}
	return SortingSpritesSceneId
}

var _ Scene = (*SortingSpritesScene)(nil)
