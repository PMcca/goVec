package vector

import (
	"math"
)

type Vector interface {
	Magnitude() float64
	Direction() float64
}

type vector2 struct {
	coordinates [2]float64
}

func (v vector2) X() float64 { return v.coordinates[0] }
func (v vector2) Y() float64 { return v.coordinates[1] }
func NewVector2(x, y float64) vector2 {
	return vector2{[2]float64{x, y}}
}

type vector3 struct {
	coordinates [3]float64
}

func (v vector3) X() float64 { return v.coordinates[0] }
func (v vector3) Y() float64 { return v.coordinates[1] }
func (v vector3) Z() float64 { return v.coordinates[2] }
func NewVector3(x, y, z float64) vector3 {
	return vector3{[3]float64{x, y, z}}
}

// 2D
func (v vector2) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X(), 2) + math.Pow(v.Y(), 2))
}

func (v vector2) Direction() float64 {
	return math.Atan(v.Y() / v.X())
}

func (v vector2) Norm() vector2 { //TODO: TEST
	return v.Div(v.Magnitude())
}

func (v vector2) Div(n float64) vector2 { //TODO TEST
	return vector2{[2]float64{v.X() / n, v.Y() / n}}
}

func (v1 vector2) Add(v2 vector2) vector2 {
	return vector2{[2]float64{v1.X() + v2.X(), v1.Y() + v2.Y()}}
}

func (v1 vector2) Sub(v2 vector2) vector2 {
	return vector2{[2]float64{v1.X() - v2.X(), v1.Y() - v2.Y()}}
}

func (v1 vector2) Dot(v2 vector2) float64 {
	return (v1.X() * v2.X()) + (v1.Y() * v2.Y())
}

// 3D
func (v vector3) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X(), 2) + math.Pow(v.Y(), 2) + math.Pow(v.Z(), 2))
}

func (v vector3) Norm() vector3 { //TODO: TEST
	return v.Div(v.Magnitude())
}

func (v vector3) Div(n float64) vector3 { //TODO TEST
	return vector3{[3]float64{v.X() / n, v.Y() / n, v.Z() / n}}
}

func (v1 vector3) Dot(v2 vector3) float64 {
	return (v1.X() * v2.X()) + (v1.Y() * v2.Y()) + (v1.Z() * v2.Z())
}

func radToDeg(r float64) float64 {
	return r * (180 / math.Pi)
}
