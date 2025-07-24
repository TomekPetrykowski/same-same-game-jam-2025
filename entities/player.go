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
	FacingUp   bool
	Vel        v.Vec
	Animations map[State]*animations.Animation
}

func (p *Player) ActiveAnimation() *animations.Animation {
	if p.FacingUp {
		return p.Animations[Up]
	} else {
		return p.Animations[Down]
	}
	// return nil
}
