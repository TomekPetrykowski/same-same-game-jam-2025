package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	Collider CollidingType
	Sprite   *Sprite
}

func NewEntity(collider CollidingType, sprite *Sprite) *Entity {
	return &Entity{
		Collider: collider,
		Sprite:   sprite,
	}
}

func (e *Entity) Update(scene Scene) {
	// print("enitityUpdate")
}

func (e *Entity) Draw(screen *ebiten.Image) {

	if e.Sprite == nil {
		DrawCollider(e.Collider, screen)

	} else {
		opts := ebiten.DrawImageOptions{}
		opts.GeoM.Translate(e.Collider.GetPos().Unpack())
		screen.DrawImage(e.Sprite.Img, &opts)
	}

}

func (e *Entity) GetCollider() CollidingType {
	return e.Collider
}
