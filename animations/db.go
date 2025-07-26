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

type AnimationData struct {
	First, Last, Step int
	Speed             float32
}

var DB = map[AnimationId]*AnimationData{
	PlayerUpId:     {First: 3, Last: 5, Step: 2, Speed: 10.0},
	PlayerDownId:   {First: 2, Last: 4, Step: 2, Speed: 10.0},
	PlayerIdleId:   {First: 0, Last: 0, Step: 0, Speed: 0.0},
  BombheadUpId:   {First: 3, Last: 5, Step: 2, Speed: 10.0},
  BombheadDownId: {First: 2, Last: 4, Step: 2, Speed: 10.0},
  BombheadIdleId: {First: 0, Last: 0, Step: 0, Speed: 0.0},
  BirdmanUpId:    {First: 0, Last: 2, Step: 2, Speed: 10.0},
  BirdmanDownId:  {First: 1, Last: 3, Step: 2, Speed: 10.0},
  BirdmanIdleId:  {First: 0, Last: 0, Step: 0, Speed: 0.0},
}
