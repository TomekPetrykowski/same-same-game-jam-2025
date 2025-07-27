package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneObjectId uint
type EnemyTypeId uint
type ObjectsMap map[SceneObjectId][]GameObject

const (
	PlayerObjectId SceneObjectId = iota
	EnemiesObjectId
	EnemyProjectilesObjectId
	PlayerProjectilesObjectId
	StaticsObjectId
	PortalsObjectId
)

const (
	BombheadTypeId EnemyTypeId = iota
	Zombiehead
	Shooter
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
