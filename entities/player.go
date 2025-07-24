package entities

import (
	"game/animations"
	v "game/utils/math"

	"github.com/hajimehoshi/ebiten/v2"
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

func (p *Player) Draw(screen *ebiten.Image) {
	p.Entity.Draw(screen)
}

func (p *Player) Update(scene Scene) {

	vel := v.Vec{}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		vel.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		vel.X += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		vel.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		vel.Y += 1
	}
	p.Collider.GetPos().Add(vel)

	for _, o := range (*scene.GetObjects())["staticObjects"] {
		p.Collider.CollideAndSlide(*o.GetCollider())
	}

}

func (p *Player) GetImage() {
	p.ActiveAnimation().Update()
	p.ActiveAnimation().Frame()

	// spritesheet := NewSpriteSheet(2, 3, 15, 26)
	// spritesheet.Rect()
}
