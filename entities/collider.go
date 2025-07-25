package entities

import (
	v "game/utils/math"
	"math"
)

type CollidingType interface {
	CollidesWith(CollidingType) bool //Checks if a the objects ovelpas with the given object
	CollideAndSlide(CollidingType)
	GetPos() *v.Vec
	SetPos(v.Vec)
}

type Rect struct {
	Pos v.Vec
	X   float64
	Y   float64
}

type Circle struct {
	Pos    v.Vec
	Radius float64
}

func (c *Circle) GetPos() *v.Vec {
	return &c.Pos
}

func (r *Rect) GetPos() *v.Vec {
	return &r.Pos
}

func (c *Circle) SetPos(v v.Vec) {
	c.Pos.X = v.X
	c.Pos.Y = v.Y
}

func (r *Rect) SetPos(v v.Vec) {
	r.Pos.X = v.X
	r.Pos.Y = v.Y
}

func (c *Circle) CollidesWith(ct CollidingType) bool {
	c2, ok := ct.(*Circle)
	if ok {
		return c.CollidesWithCircle(*c2)
	}
	r, ok := ct.(*Rect)
	if ok {
		return c.CollidesWithRect(*r)
	}
	return false
}

func (r Rect) CollidesWith(ct CollidingType) bool {
	c, ok := ct.(*Circle)
	if ok {
		return r.CollidesWithCircle(*c)
	}
	r2, ok := ct.(*Rect)
	if ok {
		return r.CollidesWithRect(*r2)
	}
	return false
}

func (c *Circle) CollidesWithRect(r Rect) bool {
	x := math.Max(r.Pos.X, math.Min(c.Pos.X, r.Pos.X+r.X))
	y := math.Max(r.Pos.Y, math.Min(c.Pos.Y, r.Pos.Y+r.Y))
	if c.Pos.DistanceTo(v.Vec{X: x, Y: y}) < c.Radius {
		return true
	}
	return false
}

func (c Circle) CollidesWithCircle(c2 Circle) bool {
	if c.Pos.DistanceTo(c2.Pos) < c.Radius+c2.Radius {
		return true
	}
	return false
}

func (r Rect) CollidesWithRect(r2 Rect) bool {
	if r.Pos.X+r.X >= r2.Pos.X &&
		r.Pos.X <= r2.Pos.X+r2.X &&
		r.Pos.Y+r.Y >= r2.Pos.Y &&
		r.Pos.Y <= r2.Pos.Y+r2.Y {
		return true
	}
	return false
}

func (r Rect) CollidesWithCircle(c Circle) bool {
	x := math.Max(r.Pos.X, math.Min(c.Pos.X, r.Pos.X+r.X))
	y := math.Max(r.Pos.Y, math.Min(c.Pos.Y, r.Pos.Y+r.Y))
	if c.Pos.DistanceTo(v.Vec{X: x, Y: y}) < c.Radius {
		return true
	}
	return false
}

func (r *Rect) CollideAndSlide(ct CollidingType) {
	c2, ok := ct.(*Circle)
	if ok {
		r.CollideAndSlideCircle(*c2)
	}
	r2, ok := ct.(*Rect)
	if ok {
		r.CollideAndSlideRect(*r2)
	}

}

func (r *Rect) CollideAndSlideRect(r2 Rect) {
	left := r2.Pos.X - (r.Pos.X + r.X)
	right := r.Pos.X - (r2.Pos.X + r2.X)
	up := r2.Pos.Y - (r.Pos.Y + r.Y)
	down := r.Pos.Y - (r2.Pos.Y + r2.Y)
	if left < 0 && right < 0 && up < 0 && down < 0 {
		if math.Max(left, right) > math.Max(up, down) {
			if left > right {
				r.Pos.AddX(left)
			} else {
				r.Pos.AddX(-right)
			}
		} else {
			if up > down {
				r.Pos.AddY(up)
			} else {
				r.Pos.AddY(-down)
			}

		}
	}

}

func (r *Rect) CollideAndSlideCircle(c Circle) {
	x := math.Max(r.Pos.X, math.Min(c.Pos.X, r.Pos.X+r.X))
	y := math.Max(r.Pos.Y, math.Min(c.Pos.Y, r.Pos.Y+r.Y))
	pos2 := v.Vec{X: x, Y: y}
	if c.Pos.DistanceTo(pos2) < c.Radius {
		norm := pos2.DirectionTo(c.Pos)
		diff := c.Radius - (c.Pos.Added(pos2.Inverted())).Length()
		r.Pos.Add(norm.Multiplied(-diff))
	}
}

func (c *Circle) CollideAndSlide(ct CollidingType) {
	c2, ok := ct.(*Circle)
	if ok {
		c.CollideAndSlideCircle(*c2)
	}
	r, ok := ct.(*Rect)
	if ok {
		c.CollideAndSlideRect(*r)
	}
}

func (c *Circle) CollideAndSlideRect(r Rect) {
	x := math.Max(r.Pos.X, math.Min(c.Pos.X, r.Pos.X+r.X))
	y := math.Max(r.Pos.Y, math.Min(c.Pos.Y, r.Pos.Y+r.Y))
	pos2 := v.Vec{X: x, Y: y}
	if c.Pos.DistanceTo(pos2) < c.Radius {
		norm := c.Pos.DirectionTo(pos2)
		diff := c.Radius - (c.Pos.Added(pos2.Inverted())).Length()
		c.Pos.Add(norm.Multiplied(-diff))
	}
}

func (c *Circle) CollideAndSlideCircle(c2 Circle) {
	if c.CollidesWithCircle(c2) {
		norm := c.Pos.DirectionTo(c2.Pos)
		diff := (c.Radius + c2.Radius) - (c.Pos.Added(c2.Pos.Inverted())).Length()
		c.Pos.Add(norm.Multiplied(-diff))
	}
}

func NewCircle(x, y, r float64) *Circle {
	return &Circle{Pos: v.Vec{X: x, Y: y}, Radius: r}
}

func NewRect(x, y, w, h float64) *Rect {
	return &Rect{Pos: v.Vec{X: x, Y: y}, X: w, Y: h}
}
