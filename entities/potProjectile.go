package entities

import (
	"game/utils/images"
	v "game/utils/math"

	"github.com/hajimehoshi/ebiten/v2"
)

type PotProjectile struct {
	*Entity
	Speed      float64
	Direction  v.Vec
	LiveLength int
	rotation   float64
}

func (p *PotProjectile) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Sprite.Offset.Unpack())
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(p.Collider.GetPos().Unpack())

	screen.DrawImage(p.Sprite.Img, &op)
	DrawCollider(p.Collider, screen)

}

func (p *PotProjectile) Update(scene Scene) {
	sceneObjects := *scene.GetObjects()
	p.rotation += 0.1

	p.LiveLength -= 1
	if p.LiveLength <= 0 {
		p.deleted = true
		return
	}

	// Move the projectile
	p.Collider.GetPos().Add(p.Direction.Multiplied(p.Speed))

	for _, obj := range sceneObjects[StaticsObjectId] {
		if p.Collider.CollidesWith(obj.GetCollider()) {
			p.deleted = true
		}
	}
	for _, obj := range sceneObjects[EnemiesObjectId] {
		if p.Collider.CollidesWith(obj.GetCollider()) {
			obj.Hit(1)
			p.deleted = true
		}
	}
}

func NewPotProjectile(x, y, speed float64, direction v.Vec) *PotProjectile {
	potImg := images.LoadImage(
		"assets/pot.png",
		"Error while loading pot image.",
		&images.DefaultPlaceholder,
	)
	ent := NewEntity(NewCircle(x, y, 3), &Sprite{Img: potImg, Offset: v.Vec{X: -5, Y: -4.5}})
	return &PotProjectile{Entity: ent, Speed: speed, Direction: direction, LiveLength: 180}

}
