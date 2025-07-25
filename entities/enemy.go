package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type BasicEnemy struct {
	Entity
	// Speed     float64
	// Direction v.Vec
}

func (e *BasicEnemy) Draw(screen *ebiten.Image) {
	// print("Enemy")
	e.Entity.Draw(screen)
}

func (e *BasicEnemy) Update(scene Scene) {
	// fmt.Println(e.Collider.GetPos())
	playerPos := (*(*scene.GetObjects())["player"][0].GetCollider()).GetPos()
	// fmt.Println(e.Collider.GetPos().DirectionTo(*playerPos))
	e.Collider.GetPos().Add(e.Collider.GetPos().DirectionTo(*playerPos).Multiplied(0.5))
	for _, o := range (*scene.GetObjects())["staticObjects"] {
		e.Collider.CollideAndSlide(*o.GetCollider())
	}
}

func NewBasicEnemy(x, y float64) *BasicEnemy {
	ent := NewEntity(NewCircle(x, y, 10))
	return &BasicEnemy{Entity: *ent}
}
