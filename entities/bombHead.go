package entities

import (
	anim "game/animations"
	"game/spritesheets"
	"game/utils/images"
	v "game/utils/math"

	"github.com/hajimehoshi/ebiten/v2"
)

type BombHead struct {
	*Entity
	Hp          int
	Speed       float64
	BlastRadius float64
	Fuse        int
	Exploding   bool
	// Animation   anim.Animation
	Animations map[State]*anim.Animation
}

func (e *BombHead) ActiveAnimation() *anim.Animation {
	return e.Animations[Down]

}
func (e *BombHead) Hit(damage int) {
	e.Hp -= damage
	if e.Hp <= 0 {
		e.Entity.deleted = true
	}
}

func (e *BombHead) Update(scene Scene) {
	if !e.Exploding {
		sceneObjects := *scene.GetObjects()
		player := sceneObjects[PlayerObjectId][0]
		playerPos := player.GetCollider().GetPos()
		e.ActiveAnimation().Update()
		if e.Collider.GetPos().DistanceTo(*playerPos) <= 15 {
			e.Exploding = true
		}
		e.Collider.GetPos().Add(e.Collider.GetPos().DirectionTo(*playerPos).Multiplied(e.Speed))

		for _, o := range sceneObjects[StaticsObjectId] {
			e.Collider.CollideAndSlide(o.GetCollider())
		}
	} else {
		e.Fuse -= 1
		if e.Fuse <= 0 {
			e.deleted = true
			scene.AddObject(
				EnemyProjectilesObjectId,
				NewExplosion(e.Collider.GetPos().Unpack()),
			)
		}
	}
}

func (e *BombHead) Draw(screen *ebiten.Image) {
	DrawSprite(screen, e.GetCurrentImage(), *e.Entity.Collider.GetPos(), e.Sprite.Offset)
	DrawCollider(e.Collider, screen)
}

func (e *BombHead) GetCurrentImage() *ebiten.Image {
	return images.SubImage(
		e.Entity.Sprite.Img,
		spritesheets.DB[spritesheets.BombheadSpritesheetId].Rect(
			e.ActiveAnimation().Frame(),
		),
	)
}

func NewBombHead(x, y float64) *BombHead {
	bombheadImg := images.LoadImage(
		"assets/bombhead.png",
		"Error while loading player image.",
		&images.DefaultPlaceholder,
	)

	animations := map[State]*anim.Animation{
		Up:   anim.NewAnimation(anim.DB[anim.BombheadUpId]),
		Down: anim.NewAnimation(anim.DB[anim.BombheadDownId]),
		Idle: anim.NewAnimation(anim.DB[anim.BombheadIdleId]),
	}

	sprite := Sprite{Img: bombheadImg, Offset: v.Vec{X: -8.5, Y: -22}}
	return &BombHead{Entity: NewEntity(NewCircle(x, y, 5), &sprite), Speed: 1, Animations: animations, Fuse: 30, Hp: 3}
}
