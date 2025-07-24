package scenes

import (
	"game/animations"
	"game/entities"
	e "game/entities"
	s "game/spritesheets"
	v "game/utils/math"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TestLevelScene struct {
	loaded            bool
	player            *e.Player
	playerSpriteSheet *s.SpriteSheet
}

func NewTestLevelScene() *TestLevelScene {
	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	playerSpriteSheet := s.NewSpriteSheet(2, 3, 15, 26)

	return &TestLevelScene{
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
		// rect:              false,s
	}
}

func (d *TestLevelScene) FirstLoad() {
}

func (d *TestLevelScene) IsLoaded() bool {
	return d.loaded
}

func (d *TestLevelScene) Draw(screen *ebiten.Image) {

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(d.player.Pos.X, d.player.Pos.Y)
	playerFrame := 0
	activeAnim := d.player.ActiveAnimation()
	if activeAnim != nil {
		playerFrame = activeAnim.Frame()
	}

	screen.DrawImage(
		d.player.Sprite.Img.SubImage(
			d.playerSpriteSheet.Rect(playerFrame),
		).(*ebiten.Image),
		&op,
	)

	// op.GeoM.Reset()

	ebitenutil.DebugPrint(screen, "Press E to change objects")
}

func (d *TestLevelScene) Update() SceneId {

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
	return TestLevelSceneId
}

func (d *TestLevelScene) OnEnter() {

}

func (d *TestLevelScene) OnExit() {

}

var _ Scene = (*TestLevelScene)(nil)
