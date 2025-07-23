package scenes

import (
	"game/animations"
	"game/entities"
	e "game/entities"
	s "game/spritesheets"
	"game/utils/draw"
	v "game/utils/math"

	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PlayerTestScene struct {
	loaded            bool
	player            *e.Player
	playerSpriteSheet *s.SpriteSheet
	c1                e.Circle
	c2                e.Circle
	r1                e.Rect
	r2                e.Rect
	rect              bool
}

func NewPlayerTestScene() *PlayerTestScene {
	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	playerSpriteSheet := s.NewSpriteSheet(2, 3, 15, 26)

	return &PlayerTestScene{
		loaded: false,
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
		rect:              false,
	}
}

func (d *PlayerTestScene) FirstLoad() {
	d.c1 = e.Circle{Pos: v.Vec{X: 90, Y: 80}, Radius: 16}
	d.c2 = e.Circle{Pos: v.Vec{X: 220, Y: 170}, Radius: 16}
	d.r1 = e.Rect{Pos: v.Vec{X: 74, Y: 154}, X: 32, Y: 32}
	d.r2 = e.Rect{Pos: v.Vec{X: 204, Y: 64}, X: 32, Y: 32}
}

func (d *PlayerTestScene) IsLoaded() bool {
	return d.loaded
}

func (d *PlayerTestScene) Draw(screen *ebiten.Image) {
	draw.DrawRect(d.r1, screen)
	draw.DrawRect(d.r2, screen)
	draw.DrawCircle(d.c1, screen)
	draw.DrawCircle(d.c2, screen)

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(d.player.Pos.X, d.player.Pos.Y)
	playerFrame := 0
	activeAnim := d.player.ActiveAnimation(d.player.Vel)
	if activeAnim != nil {
		playerFrame = activeAnim.Frame()
	}

	screen.DrawImage(
		d.player.Sprite.Img.SubImage(
			d.playerSpriteSheet.Rect(playerFrame),
		).(*ebiten.Image),
		&op,
	)

	op.GeoM.Reset()

	ebitenutil.DebugPrint(screen, "Press E to change objects")
}

func (d *PlayerTestScene) Update() SceneId {
	// player movement
	d.player.Vel = v.Vec{}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		d.player.Vel.X = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		d.player.Vel.X = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		d.player.Vel.Y = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		d.player.Vel.Y = 1
	}

	d.player.Pos.Add(d.player.Vel)

	activeAnim := d.player.ActiveAnimation(d.player.Vel)
	if activeAnim != nil {
		activeAnim.Update()
	}

	// circle movement

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
		if d.r1.CollidesWith(d.c2) || d.r1.CollidesWith(d.r2) {
			d.r1.Pos.Add(vel.Inverted())
		}
	} else {
		d.c1.Pos.Add(vel)
		if d.c1.CollidesWith(d.c2) || d.c1.CollidesWith(d.r2) {
			d.c1.Pos.Add(vel.Inverted())
		}
	}
	return PlayerTestSceneId
}

func (d *PlayerTestScene) OnEnter() {

}

func (d *PlayerTestScene) OnExit() {

}

var _ Scene = (*PlayerTestScene)(nil)
