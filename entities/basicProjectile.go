package entities

import (
	"fmt"
	v "game/utils/math"
)

type BasicProjectile struct {
	*Entity
	Speed      float64
	Direction  v.Vec
	LiveLength int
}

func (bp *BasicProjectile) Update(scene Scene) {
	sceneObjects := *scene.GetObjects()
	player := sceneObjects[PlayerObjectId][0]

	bp.LiveLength -= 1
	if bp.LiveLength <= 0 {
		bp.deleted = true
		return
	}

	// Move the projectile
	bp.Collider.GetPos().Add(bp.Direction.Multiplied(bp.Speed))

	if bp.Collider.CollidesWith(player.GetCollider()) {
		fmt.Println("Player hit")
		player.Hit(2)
		bp.deleted = true
		return
	}
	for _, obj := range sceneObjects[StaticsObjectId] {
		if bp.Collider.CollidesWith(obj.GetCollider()) {
			bp.deleted = true
			return
		}
	}
}

func NewBasicProjectile(x, y, speed float64, direction v.Vec) *BasicProjectile {
	ent := NewEntity(NewCircle(x, y, 3), nil)
	return &BasicProjectile{Entity: ent, Speed: speed, Direction: direction, LiveLength: 120}
}
