package scenes

// import (
// 	"game/animations"
// 	"game/entities"
// 	e "game/entities"
// 	s "game/spritesheets"
// 	v "game/utils/math"

// 	"log"

// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
// )

// type PlayerTestScene struct {
// 	loaded            bool
// 	player            *e.Player
// 	playerSpriteSheet *s.SpriteSheet
// 	c1                e.Circle
// 	c2                e.Circle
// 	r1                e.Rect
// 	r2                e.Rect
// 	rect              bool
// }

// func NewPlayerTestScene() *PlayerTestScene {
// 	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
// 	if err != nil {
// 		// handle error
// 		log.Fatal(err)
// 	}
// 	playerSpriteSheet := s.NewSpriteSheet(2, 3, 15, 26)

// 	return &PlayerTestScene{
// 		loaded: false,
// 		player: &entities.Player{
// 			Entity: entities.Entity{
// 				Sprite: &entities.Sprite{
// 					Img: playerImg,
// 					X:   50.0,
// 					Y:   50.0,
// 				},
// 				Collider: &e.Circle{Pos: v.Vec{X: 0, Y: 0}, Radius: 10},
// 			},
// 			Animations: map[entities.State]*animations.Animation{
// 				entities.Up:   animations.NewAnimation(3, 5, 2, 20.0),
// 				entities.Down: animations.NewAnimation(2, 4, 2, 20.0),
// 			},
// 		},
// 		playerSpriteSheet: playerSpriteSheet,
// 		rect:              false,
// 	}
// }

// func (d *PlayerTestScene) FirstLoad() {

// }

// func (d *PlayerTestScene) IsLoaded() bool {
// 	return d.loaded
// }

// func (d *PlayerTestScene) Draw(screen *ebiten.Image) {

// 	op := ebiten.DrawImageOptions{}
// 	op.GeoM.Translate(d.player.Pos.X, d.player.Pos.Y)
// 	playerFrame := 0
// 	activeAnim := d.player.ActiveAnimation()
// 	if activeAnim != nil {
// 		playerFrame = activeAnim.Frame()
// 	}

// 	d.player.Draw(screen)
// 	screen.DrawImage(
// 		d.player.Sprite.Img.SubImage(
// 			d.playerSpriteSheet.Rect(playerFrame),
// 		).(*ebiten.Image),
// 		&op,
// 	)

// 	op.GeoM.Reset()

// 	ebitenutil.DebugPrint(screen, "Press E to change objects")
// }

// func (d *PlayerTestScene) Update() SceneId {
// 	// player movement
// 	vel := v.Vec{}

// 	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
// 		vel.X = -1
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyRight) {
// 		vel.X = 1
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyUp) {
// 		vel.Y = -1
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyDown) {
// 		vel.Y = 1
// 	}

// 	if vel.Y < 0 {
// 		d.player.FacingUp = true
// 	} else if vel.Y > 0 {
// 		d.player.FacingUp = false
// 	}

// 	d.player.Collider.GetPos().Add(vel)

// 	activeAnim := d.player.ActiveAnimation()
// 	if activeAnim != nil {
// 		activeAnim.Update()
// 	}

// 	return PlayerTestSceneId
// }

// func (d *PlayerTestScene) OnEnter() {

// }

// func (d *PlayerTestScene) OnExit() {

// }

// var _ Scene = (*PlayerTestScene)(nil)
