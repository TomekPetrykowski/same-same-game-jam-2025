package entities

import (
	"game/animations"
	v "game/utils/math"
)

type State uint8

const (
	Down State = iota
	Up
)

type Player struct {
	Entity
	Vel        v.Vec
	Animations map[State]*animations.Animation
}

func (p *Player) ActiveAnimation(vel v.Vec) *animations.Animation {
	if vel.Y > 0 {
		return p.Animations[Down]
	}
	if vel.Y < 0 {
		return p.Animations[Up]
	}
	return nil
}
