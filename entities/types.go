package entities

import "github.com/hajimehoshi/ebiten/v2"

type SceneObjectId uint

const (
	PlayerObjectId SceneObjectId = iota
	EnemiesObjectId
	EnemyProjectilesObjectId
	PlayerProjectilesObjectId
	StaticsObjectId
)

type Scene interface {
	GetObjects() *map[SceneObjectId][]GameObject
	AddObject(SceneObjectId, GameObject)
}

type GameObject interface {
	GetCollider() CollidingType
	Update(Scene)
	Draw(*ebiten.Image)
	IsDeleted() bool
}
