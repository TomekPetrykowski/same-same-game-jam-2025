package animations

type AnimationId uint

const (
	PlayerUpId AnimationId = iota
	PlayerDownId
	PlayerIdleId
	BombHeadUpId
	BombHeadDownId
	BombHeadIdleId
)

var DB = map[AnimationId]*Animation{
	PlayerUpId:     NewAnimation(3, 5, 2, 10.0),
	PlayerDownId:   NewAnimation(2, 4, 2, 10.0),
	PlayerIdleId:   NewAnimation(0, 0, 0, 0.0),
	BombHeadUpId:   NewAnimation(3, 5, 2, 10.0),
	BombHeadDownId: NewAnimation(2, 4, 2, 10.0),
	BombHeadIdleId: NewAnimation(0, 0, 0, 0.0),
}
