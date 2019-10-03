package vector

import (
	"math"
)

type Vector interface {
	Magnitude() float64
	Direction() float64
}

type Vector2 struct {
	X float64
	Y float64
}

type Vector3 struct {
	Vector2
	Z float64
}

func (v Vector2) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func (v Vector2) Direction() float64 {
	return math.Atan(v.Y / v.X)
}

func (v1 Vector2) Dot(v2 Vector2) float64 {
	return (v1.X * v2.X) + (v1.Y * v2.Y)
}

func radToDeg(r float64) float64 {
	return r * (180 / math.Pi)
}
