package entities

import (
	anim "game/animations"
	spritesheet "game/spritesheets"
	"game/utils/images"
	v "game/utils/math"

	"github.com/hajimehoshi/ebiten/v2"
)

type State uint8

const (
	Down State = iota
	Up
	Idle
)

type Player struct {
	*Entity
	FacingUp    bool
	Vel         v.Vec
	Animations  map[State]*anim.Animation
	Spritesheet *spritesheet.Spritesheet
}

func (p *Player) ActiveAnimation() *anim.Animation {

	if p.Vel.IsZero() {
		return p.Animations[Idle]
	}

	if p.FacingUp {
		return p.Animations[Up]
	} else {
		return p.Animations[Down]
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	DrawCollider(p.Collider, screen)
	DrawSprite(screen, p.GetCurrentImage(), *p.Entity.Collider.GetPos(), v.Vec{X: -7.5, Y: -20})
}

func (p *Player) Update(scene Scene) {
	// Moved updating animation in releveant method
	p.ActiveAnimation().Update()

	p.Vel.Reset()
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Vel.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Vel.X += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Vel.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Vel.Y += 1
	}
	p.Collider.GetPos().Add(p.Vel)
	if p.Vel.Y < 0 {
		p.FacingUp = true
	} else if p.Vel.Y > 0 {
		p.FacingUp = false
	}

	sceneObjects := *scene.GetObjects()
	for _, o := range sceneObjects[StaticsObjectId] {
		p.Collider.CollideAndSlide(o.GetCollider())
	}

}

func (p *Player) GetCurrentImage() *ebiten.Image {
	return images.SubImage(
		p.Entity.Sprite.Img,
		p.Spritesheet.Rect(
			p.ActiveAnimation().Frame(),
		),
	)
}

func NewPlayer(initX, initY float64) *Player {
	playerImg := images.LoadImage(
		"assets/player.png",
		"Error while loading player image.",
		&images.DefaultPlaceholder,
	)

	animations := map[State]*anim.Animation{
		Up:   anim.DB[anim.PlayerUpId],
		Down: anim.DB[anim.PlayerDownId],
		Idle: anim.DB[anim.PlayerIdleId],
	}

	sprite := Sprite{Img: playerImg}
	ent := NewEntity(NewCircle(initX, initY, 5), &sprite)

	return &Player{
		Entity:      ent,
		Animations:  animations,
		Spritesheet: spritesheet.DB[spritesheet.PlayerSpritesheedId],
	}
}
