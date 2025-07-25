package entities

import "github.com/hajimehoshi/ebiten/v2"

type SceneObject uint

const (
	PlayerObjectId SceneObject = iota
	EnemiesObjectId
	EnemyProjectilesObjectId
	PlayerProjectilesObjectId
	StaticsObjectId
)

type Scene interface {
	GetObjects() *map[SceneObject][]GameObject
  AddObject(SceneObject, GameObject)
}

type GameObject interface {
	GetCollider() *CollidingType
	Update(Scene)
	Draw(*ebiten.Image)
}
