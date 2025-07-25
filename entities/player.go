package entities

import (
	"game/animations"
	s "game/spritesheets"
	v "game/utils/math"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	DrawCollider(p.Collider, screen)
	DrawSprite(screen, p.GetImage(), *p.Entity.Collider.GetPos(), v.Vec{X: -7.5, Y: -20})
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
	if vel.Y < 0 {
		p.FacingUp = true
	} else if vel.Y > 0 {
		p.FacingUp = false
	}

	for _, o := range (*scene.GetObjects())["staticObjects"] {
		p.Collider.CollideAndSlide(*o.GetCollider())
	}

}

func (p *Player) GetImage() *ebiten.Image {

	p.ActiveAnimation().Update()

	return (p.Entity.Sprite.Img.SubImage(s.DB[s.PlayerSpritesheedId].Rect(p.ActiveAnimation().Frame()))).(*ebiten.Image)

}

func NewPlayer(x, y float64) *Player {
	sprite, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	animations := map[State]*animations.Animation{
		Up:   animations.NewAnimation(3, 5, 2, 20.0),
		Down: animations.NewAnimation(2, 4, 2, 20.0),
	}
	ent := *NewEntity(NewCircle(x, y, 5), &Sprite{Img: sprite})
	// img := ebiten.NewImage(20, 20)
	// img.Fill(color.White)
	// ent.Sprite.Img = img
	return &Player{Entity: ent, Animations: animations}
}
