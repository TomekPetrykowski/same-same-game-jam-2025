package animations

type Animation struct {
	First        int
	Last         int
	Step         int
	SpeedInTps   float32
	frameCounter float32
	frame        int
}

func (a *Animation) Update() {
	a.frameCounter -= 1.0
	if a.frameCounter < 0.0 {
		a.frameCounter = a.SpeedInTps
		a.frame += a.Step
		if a.frame > a.Last {
			// loop back to the beginning
			a.frame = a.First
		}
	}
}

func (a *Animation) Frame() int {
	return a.frame
}

func NewAnimation(data *AnimationData) *Animation {
	return &Animation{
		data.First,
		data.Last,
		data.Step,
		data.Speed,
		data.Speed,
		data.First,
	}
}

func (a *Animation) Reset() {
	a.frameCounter = a.SpeedInTps
	a.frame = a.First
}
