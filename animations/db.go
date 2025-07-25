package animations

type AnimationId uint

const (
	PlayerUpId AnimationId = iota
	PlayerDownId
	PlayerIdleId
	BombheadUpId
	BombheadDownId
	BombheadIdleId
	BirdmanUpId
	BirdmanDownId
	BirdmanIdleId
)

var DB = map[AnimationId]*Animation{
	PlayerUpId:     NewAnimation(3, 5, 2, 10.0),
	PlayerDownId:   NewAnimation(2, 4, 2, 10.0),
	PlayerIdleId:   NewAnimation(0, 0, 0, 0.0),
	BombheadUpId:   NewAnimation(3, 5, 2, 10.0),
	BombheadDownId: NewAnimation(2, 4, 2, 10.0),
	BombheadIdleId: NewAnimation(0, 0, 0, 0.0),
	BirdmanUpId:    NewAnimation(0, 2, 2, 10.0),
	BirdmanDownId:  NewAnimation(1, 3, 2, 10.0),
	BirdmanIdleId:  NewAnimation(0, 0, 0, 0.0),
}
