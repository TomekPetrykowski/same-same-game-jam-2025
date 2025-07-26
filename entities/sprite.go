package entities

import (
	v "game/utils/math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Offset v.Vec
	Img    *ebiten.Image
}
