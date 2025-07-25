package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type BasicEnemy struct {
	*Entity
	Speed float64
	// Direction v.Vec
}

func (be *BasicEnemy) Draw(screen *ebiten.Image) {
	be.Entity.Draw(screen)
}

func (be *BasicEnemy) Update(scene Scene) {

	sceneObjects := *scene.GetObjects()
	player := sceneObjects[PlayerObjectId][0]
	playerPos := player.GetCollider().GetPos()

	// setting enemy direction and movement towards player
	be.Collider.GetPos().Add(be.Collider.GetPos().DirectionTo(*playerPos).Multiplied(be.Speed))

	// setting colliders with static objects on scene
	for _, o := range sceneObjects[StaticsObjectId] {
		be.Collider.CollideAndSlide(o.GetCollider())
	}
}

func NewBasicEnemy(x, y, speed float64) *BasicEnemy {
	ent := NewEntity(NewCircle(x, y, 10), nil)
	return &BasicEnemy{Entity: ent, Speed: speed}
}
