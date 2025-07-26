package entities

import (
	"fmt"
	"game/utils/images"
	v "game/utils/math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Explosion struct {
	*Entity
	Speed      float64
	Direction  v.Vec
	LiveLength int
}

func (e *Explosion) Draw(screen *ebiten.Image) {
	DrawSprite(screen, e.Sprite.Img, *e.Collider.GetPos(), e.Sprite.Offset)
}

func (e *Explosion) Update(scene Scene) {
	sceneObjects := *scene.GetObjects()
	player := sceneObjects[PlayerObjectId][0]

	e.LiveLength -= 1
	if e.LiveLength <= 0 {
		e.deleted = true
	}

	if e.Collider.CollidesWith(player.GetCollider()) {
		// print("player hit") // print outputs to standard error not to standard output
		fmt.Println("Player hit")
	}
}

func NewExplosion(x, y float64) *Explosion {
	img := images.LoadImage(
		"assets/sun.png",
		"Error while loading player image.",
		&images.DefaultPlaceholder,
	)
	ent := NewEntity(NewCircle(x, y, 5), &Sprite{Img: img, Offset: v.Vec{X: -25, Y: -24.5}})
	return &Explosion{Entity: ent, LiveLength: 60}

}
