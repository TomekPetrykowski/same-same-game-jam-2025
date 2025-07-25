package entities

import (
	v "game/utils/math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	Pos      v.Vec
	Collider CollidingType
	Sprite   *Sprite
}

type Scene interface {
	GetObjects() *map[string][]GameObject
	AddObject(string, GameObject)
}

func (e *Entity) Update(scene Scene) {
	// print("enitityUpdate")
}

type GameObject interface {
	GetCollider() *CollidingType
	Update(Scene)
	Draw(*ebiten.Image)
}

func (e *Entity) Draw(screen *ebiten.Image) {

	if e.Sprite == nil {
		DrawCollider(e.Collider, screen)

	} else {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(e.Collider.GetPos().Unpack())
		screen.DrawImage(e.Sprite.Img, opts)
	}

}

func (e *Entity) GetCollider() *CollidingType {
	return &e.Collider
}

func NewEntity(ct CollidingType) *Entity {
	return &Entity{Collider: ct}
}
