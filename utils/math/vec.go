package utils

import (
	"math"
	"strconv"
)

type Vec struct {
	X float64
	Y float64
}

func (v Vec) Added(v2 Vec) Vec {
	return Vec{v.X + v2.X, v.Y + v2.Y}
}

func (v *Vec) Add(v2 Vec) {
	v.X += v2.X
	v.Y += v2.Y
}

func (v *Vec) Invert() {
	v.X = -v.X
	v.Y = -v.Y
}

func (v Vec) Inverted() Vec {
	return Vec{-v.X, -v.Y}
}

func (v Vec) String() string {
	return "(X:" + strconv.FormatFloat(v.X, 'f', 2, 64) + ",Y:" + strconv.FormatFloat(v.Y, 'f', 2, 64) + ")"
}

func (v *Vec) Normalize() {
	l := v.Length()
	v.X /= l
	v.Y /= l
}

func (v Vec) Normalized() Vec {
	l := v.Length()
	return Vec{v.X / l, v.Y / l}
}

func (v Vec) DistanceTo(v2 Vec) float64 {
	return Vec{v.X - v2.X, v.Y - v2.Y}.Length()
}

func (v Vec) Length() float64 {
	return (math.Sqrt((v.X * v.X) + (v.Y * v.Y)))
}

func (v Vec) Angle() float64 {
	return math.Asin(v.X / v.Length())
}

func (v Vec) DirectionTo(v2 Vec) Vec {
	return (v.Added(v2.Inverted())).Normalized()
}

func (v Vec) Times(a float64) Vec {
	return Vec{v.X * a, v.Y * a}
}

func (v *Vec) Multiply(a float64) {
	v.X *= a
	v.Y *= a
}
func (v Vec) Multiplied(a float64) Vec {
	return Vec{v.X * a, v.Y * a}
}

func (v Vec) Equals(v2 Vec) bool {
	return ((v.X == v2.X) && (v.Y == v2.Y))
}
