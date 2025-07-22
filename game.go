package main

import (
	"fmt"
	"game/animations"
	"game/entities"
	e "game/entities"
	s "game/spritesheets"
	v "game/utils/math"

	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var ballSprite, _, _ = ebitenutil.NewImageFromFile("assets/ball.png")

type Game struct {
	player            *e.Player
	playerSpriteSheet *s.SpriteSheet
	c1                e.Circle
	c2                e.Circle
	r1                e.Rect
	r2                e.Rect
	rect              bool
}

func NewGame() *Game {
	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	playerSpriteSheet := s.NewSpriteSheet(2, 3, 15, 26)

	return &Game{
		player: &entities.Player{
			Entity: entities.Entity{
				Sprite: &entities.Sprite{
					Img: playerImg,
					X:   50.0,
					Y:   50.0,
				},
			},
			Animations: map[entities.State]*animations.Animation{
				entities.Up:   animations.NewAnimation(3, 5, 2, 20.0),
				entities.Down: animations.NewAnimation(2, 4, 2, 20.0),
			},
		},
		playerSpriteSheet: playerSpriteSheet,
	}
}

func (g *Game) Update() error {
	g.player.Vel = v.Vec{}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.Vel.X = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.Vel.X = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.player.Vel.Y = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.player.Vel.Y = 1
	}

	g.player.Pos.Add(g.player.Vel)

	activeAnim := g.player.ActiveAnimation(g.player.Vel)
	if activeAnim != nil {
		activeAnim.Update()
	}

	vel := v.Vec{}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		g.rect = !g.rect
	}
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
	if g.rect {
		g.r1.Pos.Add(vel)
		if g.r1.CollidesWith(g.c2) || g.r1.CollidesWith(g.r2) {
			g.r1.Pos.Add(vel.Inverted())
		}
	} else {
		g.c1.Pos.Add(vel)
		if g.c1.CollidesWith(g.c2) || g.c1.CollidesWith(g.r2) {
			g.c1.Pos.Add(vel.Inverted())
		}
	}
	return nil
}

func DrawRect(r e.Rect, screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	img := ebiten.NewImage(int(r.X), int(r.Y))
	img.Fill(color.White)
	op.GeoM.Translate(r.Pos.X, r.Pos.Y)
	screen.DrawImage(img, &op)
}

func DrawCircle(c e.Circle, screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	img := ebiten.NewImageFromImage(ballSprite)
	scale := float64(float64(c.Radius) * 2 / 32)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(-(float64(c.Radius)), -(float64(c.Radius)))
	op.GeoM.Translate(c.Pos.X, c.Pos.Y)
	screen.DrawImage(img, &op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawRect(g.r1, screen)
	DrawRect(g.r2, screen)
	DrawCircle(g.c1, screen)
	DrawCircle(g.c2, screen)

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.player.Pos.X, g.player.Pos.Y)
	playerFrame := 0
	activeAnim := g.player.ActiveAnimation(g.player.Vel)
	if activeAnim != nil {
		playerFrame = activeAnim.Frame()
	}

	screen.DrawImage(
		g.player.Sprite.Img.SubImage(
			g.playerSpriteSheet.Rect(playerFrame),
		).(*ebiten.Image),
		&op,
	)

	op.GeoM.Reset()

	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) Set() {

	g.c1 = e.Circle{Pos: v.Vec{X: 90, Y: 80}, Radius: 16}
	g.c2 = e.Circle{Pos: v.Vec{X: 220, Y: 170}, Radius: 16}
	g.r1 = e.Rect{Pos: v.Vec{X: 74, Y: 154}, X: 32, Y: 32}
	g.r2 = e.Rect{Pos: v.Vec{X: 204, Y: 64}, X: 32, Y: 32}

	fmt.Println("Game is running!")
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
}
