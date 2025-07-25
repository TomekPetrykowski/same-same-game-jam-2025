package entities

import (
	v "game/utils/math"
)

type BasicProjectile struct {
	Entity
	Speed      float64
	Direction  v.Vec
	LiveLength int
}

func (p BasicProjectile) Update(scene Scene) {
	p.LiveLength -= 1
	if p.LiveLength <= 0 {
		//destroy
	}
	p.Collider.GetPos().Add(p.Direction.Multiplied(p.Speed))
	if p.Collider.CollidesWith(*(*scene.GetObjects())["player"][0].GetCollider()) {
		print("player hit")
	}
	for _, obj := range (*scene.GetObjects())["staticObjects"] {
		if p.Collider.CollidesWith(*obj.GetCollider()) {
			print("Destroyed")
		}
	}

}

func NewBasicProjectile(x, y, speed float64, direction v.Vec) *BasicProjectile {
	ent := NewEntity(NewCircle(x, y, 3))
	return &BasicProjectile{Entity: *ent, Speed: speed, Direction: direction, LiveLength: 120}

}
