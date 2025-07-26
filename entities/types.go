package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneObjectId uint
type ObjectsMap map[SceneObjectId][]GameObject

const (
	PlayerObjectId SceneObjectId = iota
	EnemiesObjectId
	EnemyProjectilesObjectId
	PlayerProjectilesObjectId
	StaticsObjectId
)

type Scene interface {
	GetObjects() *ObjectsMap
	AddObject(SceneObjectId, GameObject)
}

type GameObject interface {
	GetCollider() CollidingType
	Update(Scene)
	Draw(*ebiten.Image)
	IsDeleted() bool
	Hit(damage int)
}
