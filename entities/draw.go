package entities

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var ballSprite, _, _ = ebitenutil.NewImageFromFile("assets/ball.png")

func DrawRect(r Rect, screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	img := ebiten.NewImage(int(r.X), int(r.Y))
	img.Fill(color.White)
	op.GeoM.Translate(r.Pos.X, r.Pos.Y)
	screen.DrawImage(img, &op)
}

func DrawCircle(c Circle, screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	img := ebiten.NewImageFromImage(ballSprite)
	scale := float64(float64(c.Radius) * 2 / 32)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(-(float64(c.Radius)), -(float64(c.Radius)))
	op.GeoM.Translate(c.Pos.X, c.Pos.Y)
	screen.DrawImage(img, &op)
}

func DrawCollider(ct CollidingType, screen *ebiten.Image) {
	r, ok := ct.(*Rect)
	if ok {
		DrawRect(*r, screen)
	}
	c, ok := ct.(*Circle)
	if ok {
		DrawCircle(*c, screen)
	}

}
