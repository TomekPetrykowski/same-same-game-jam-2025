package entities

import (
	"fmt"
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
	FacingUp           bool
	Vel                v.Vec
	Animations         map[State]*anim.Animation
	Spritesheet        *spritesheet.Spritesheet
	ShootDelay         int
	InvisibilityFrames int
	Hp                 int
}

func (p *Player) Hit(damage int) {
	if p.InvisibilityFrames <= 0 {
		p.Hp -= damage
		p.InvisibilityFrames = 60
		if p.Hp <= 0 {
			fmt.Println("player died")
			//game over
		}
	}
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
	if p.InvisibilityFrames <= 0 || p.InvisibilityFrames%2 == 0 {
		DrawSprite(screen, p.GetCurrentImage(), *p.Entity.Collider.GetPos(), p.Sprite.Offset)
	}

}

func (p *Player) Update(scene Scene) {
	// Moved updating animation in releveant method
	p.ActiveAnimation().Update()
	if p.InvisibilityFrames >= 0 {
		p.InvisibilityFrames -= 1
	}

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
	p.Vel.Normalize()
	p.Collider.GetPos().Add(p.Vel)
	if p.Vel.Y < 0 {
		p.FacingUp = true
	} else if p.Vel.Y > 0 {
		p.FacingUp = false
	}
	if p.ShootDelay > 0 {
		p.ShootDelay -= 1
	}
	if p.ShootDelay <= 0 && ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		p.ShootDelay = 30
		x, y := ebiten.CursorPosition()
		mousePos := v.Vec{X: float64(x), Y: float64(y)}
		direction := p.Collider.GetPos().DirectionTo(mousePos)
		scene.AddObject(PlayerProjectilesObjectId, NewPotProjectile(p.Collider.GetPos().X, p.Collider.GetPos().Y, 1.5, direction))
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
		Up:   anim.NewAnimation(anim.DB[anim.PlayerUpId]),
		Down: anim.NewAnimation(anim.DB[anim.PlayerDownId]),
		Idle: anim.NewAnimation(anim.DB[anim.PlayerIdleId]),
	}

	sprite := Sprite{Img: playerImg, Offset: v.Vec{X: -7.5, Y: -20}}
	ent := NewEntity(NewCircle(initX, initY, 5), &sprite)

	return &Player{
		Entity:      ent,
		Animations:  animations,
		Spritesheet: spritesheet.DB[spritesheet.PlayerSpritesheetId],
		Hp:          3,
	}
}
