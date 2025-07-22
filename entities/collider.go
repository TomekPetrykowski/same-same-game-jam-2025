package entities

import (
	v "game/utils/math"
	"math"
)

type CollidingType interface {
	CollidesWith(CollidingType) bool
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

func (c Circle) CollidesWith(ct CollidingType) bool {
	c2, ok := ct.(Circle)
	if ok {
		return c.CollidesWithCircle(c2)
	}
	r, ok := ct.(Rect)
	if ok {
		return c.CollidesWithRect(r)
	}
	return false
}

func (r Rect) CollidesWith(ct CollidingType) bool {
	c, ok := ct.(Circle)
	if ok {
		return r.CollidesWithCircle(c)
	}
	r2, ok := ct.(Rect)
	if ok {
		return r.CollidesWithRect(r2)
	}
	return false
}

func (c Circle) CollidesWithRect(r Rect) bool {
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
