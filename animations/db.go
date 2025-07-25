package animations

type AnimationId uint

const (
	PlayerUpId AnimationId = iota
	PlayerDownId
	PlayerIdleId
)

type AnimationData struct {
	First, Last, Step int
	Speed             float32
}

var DB = map[AnimationId]*AnimationData{
	PlayerUpId:   {First: 3, Last: 5, Step: 2, Speed: 10.0},
	PlayerDownId: {First: 2, Last: 4, Step: 2, Speed: 10.0},
	PlayerIdleId: {First: 0, Last: 0, Step: 0, Speed: 0.0},
}
