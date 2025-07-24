package scenes

import (
	e "game/entities"
	"game/utils/draw"
	v "game/utils/math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type DebugScene struct {
	loaded bool
	c1     e.Circle
	c2     e.Circle
	r1     e.Rect
	r2     e.Rect
	rect   bool
}

func NewDebugScene() *DebugScene {
	return &DebugScene{
		loaded: false,
		rect:   false,
	}
}

func (d *DebugScene) FirstLoad() {
	d.c1 = e.Circle{Pos: v.Vec{X: 90, Y: 80}, Radius: 16}
	d.c2 = e.Circle{Pos: v.Vec{X: 220, Y: 170}, Radius: 16}
	d.r1 = e.Rect{Pos: v.Vec{X: 74, Y: 154}, X: 32, Y: 32}
	d.r2 = e.Rect{Pos: v.Vec{X: 204, Y: 64}, X: 32, Y: 32}
}

func (d *DebugScene) IsLoaded() bool {
	return d.loaded
}

func (d *DebugScene) Draw(screen *ebiten.Image) {
	draw.DrawRect(d.r1, screen)
	draw.DrawRect(d.r2, screen)
	draw.DrawCircle(d.c1, screen)
	draw.DrawCircle(d.c2, screen)

	ebitenutil.DebugPrint(screen, "Press E to change objects")
}

func (d *DebugScene) Update() SceneId {
	vel := v.Vec{}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		d.rect = !d.rect
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
	if d.rect {
		d.r1.Pos.Add(vel)
		if d.r1.CollidesWith(&d.c2) || d.r1.CollidesWith(&d.r2) {
			d.r1.Pos.Add(vel.Inverted())
		}
	} else {
		// d.c1.Pos.Add(vel)
		d.c1.GetPos().Add(vel)
		d.c1.CollideAndSlide(&d.c2)
		d.c1.CollideAndSlide(&d.r2)
	}
	return DebugSceneId
}

func (d *DebugScene) OnEnter() {

}

func (d *DebugScene) OnExit() {

}

var _ Scene = (*DebugScene)(nil)
