package entities

import (
	v "game/utils/math"
)

type Entity struct {
	Pos      v.Vec
	Collider CollidingType
	Sprite   *Sprite
}
