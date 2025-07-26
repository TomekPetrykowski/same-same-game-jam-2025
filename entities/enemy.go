package entities

import (
	anim "game/animations"
	spritesheet "game/spritesheets"
	"game/utils/images"
	v "game/utils/math"

	"github.com/hajimehoshi/ebiten/v2"
)

type BasicEnemy struct {
	*Entity
	Speed       float64
	FacingUp    bool
	Vel         v.Vec
	Animations  map[State]*anim.Animation
	Spritesheet *spritesheet.Spritesheet
}

func (be *BasicEnemy) Draw(screen *ebiten.Image) {

	DrawCollider(be.Collider, screen)
	DrawSprite(screen, be.GetCurrentImage(), *be.Entity.Collider.GetPos(), v.Vec{X: -9, Y: -20})

}

func (be *BasicEnemy) Update(scene Scene) {
	be.ActiveAnimation().Update()

	sceneObjects := *scene.GetObjects()
	player := sceneObjects[PlayerObjectId][0]
	playerPos := player.GetCollider().GetPos()

	// setting enemy direction and movement towards player

	direction := be.Collider.GetPos().DirectionTo(*playerPos)
	be.Vel = direction.Multiplied(be.Speed)
	be.Collider.GetPos().Add(be.Vel)
	be.FacingUp = be.Vel.Y < 0

	// setting colliders with static objects on scene
	for _, o := range sceneObjects[StaticsObjectId] {
		be.Collider.CollideAndSlide(o.GetCollider())
	}
}

func (be *BasicEnemy) ActiveAnimation() *anim.Animation {

	if be.Vel.IsZero() {
		return be.Animations[Idle]
	}

	if be.FacingUp {
		return be.Animations[Up]
	} else {
		return be.Animations[Down]
	}
}

func (be *BasicEnemy) GetCurrentImage() *ebiten.Image {
	frame := be.ActiveAnimation().Frame()
	rect := be.Spritesheet.Rect(frame)

	return images.SubImage(
		be.Entity.Sprite.Img,
		rect,
	)
}

func NewBasicEnemy(x, y, speed float64) *BasicEnemy {
	enemyImg := images.LoadImage(
		"assets/bombhead.png",
		"Error while loading enemy image.",
		&images.DefaultPlaceholder,
	)

	animations := map[State]*anim.Animation{
		Up:   anim.NewAnimation(anim.DB[anim.BombheadUpId]),
		Down: anim.NewAnimation(anim.DB[anim.BombheadDownId]),
		Idle: anim.NewAnimation(anim.DB[anim.BombheadIdleId]),
	}

	sprite := Sprite{Img: enemyImg}
	ent := NewEntity(NewCircle(x, y, 5), &sprite)

	return &BasicEnemy{
		Entity:      ent,
		Speed:       speed,
		Animations:  animations,
		Spritesheet: spritesheet.DB[spritesheet.BombheadSpritesheedId],
	}
}
