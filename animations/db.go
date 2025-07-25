package animations

type AnimationId uint

const (
	PlayerUpId AnimationId = iota
	PlayerDownId
)

var DB = map[AnimationId]*Animation{
	PlayerUpId:   NewAnimation(3, 5, 2, 20.0),
	PlayerDownId: NewAnimation(2, 4, 2, 20.0),
}
