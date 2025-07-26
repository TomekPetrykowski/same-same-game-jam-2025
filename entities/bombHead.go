package entities

import (
	"game/animations"
	anim "game/animations"
	"game/spritesheets"
	"game/utils/images"
	v "game/utils/math"

	"github.com/hajimehoshi/ebiten/v2"
)

type BombHead struct {
	Entity
	Hp          float64
	Speed       float64
	BlastRadius float64
	Fuse        float64
	Exploding   bool
	Animation   anim.Animation
}

func (e *BombHead) ActiveAnimation() *anim.Animation {
	return animations.DB[animations.BombHeadDownId]
}

func (e *BombHead) Update(scene Scene) {

	sceneObjects := *scene.GetObjects()
	player := sceneObjects[PlayerObjectId][0]
	playerPos := player.GetCollider().GetPos()
	e.ActiveAnimation().Update()

	// setting enemy direction and movement towards player
	e.Collider.GetPos().Add(e.Collider.GetPos().DirectionTo(*playerPos).Multiplied(e.Speed))

	// setting colliders with static objects on scene
	for _, o := range sceneObjects[StaticsObjectId] {
		e.Collider.CollideAndSlide(o.GetCollider())
	}
}

func NewBombHead(x, y float64) *BombHead {
	bombheadImg := images.LoadImage(
		"assets/bombhead.png",
		"Error while loading player image.",
		&images.DefaultPlaceholder,
	)
	sprite := Sprite{Img: bombheadImg, Offset: v.Vec{X: -8.5, Y: -22}}
	return &BombHead{Entity: *NewEntity(NewCircle(x, y, 5), &sprite)}
}

func (e *BombHead) Draw(screen *ebiten.Image) {
	DrawCollider(e.Collider, screen)
	DrawSprite(screen, e.GetCurrentImage(), *e.Entity.Collider.GetPos(), e.Sprite.Offset)
}

func (e *BombHead) GetCurrentImage() *ebiten.Image {
	return images.SubImage(
		e.Entity.Sprite.Img,
		spritesheets.DB[spritesheets.BombheadSpritesheetId].Rect(
			e.ActiveAnimation().Frame(),
		),
	)
}
