package vector

import (
	"math"
)

type Vector interface {
	Magnitude() float64
	Norm() Vector
	X() float64
	Y() float64
}

type Vector2 struct {
	coordinates [2]float64
}

func (v Vector2) X() float64 { return v.coordinates[0] }
func (v Vector2) Y() float64 { return v.coordinates[1] }

func NewVector2(x, y float64) Vector2 {
	return Vector2{[2]float64{x, y}}
}

type Vector3 struct {
	coordinates [3]float64
}

func (v Vector3) X() float64 { return v.coordinates[0] }
func (v Vector3) Y() float64 { return v.coordinates[1] }
func (v Vector3) Z() float64 { return v.coordinates[2] }

func NewVector3(x, y, z float64) Vector3 {
	return Vector3{[3]float64{x, y, z}}
}

// 2D
func (v Vector2) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X(), 2) + math.Pow(v.Y(), 2))
}

func (v Vector2) Direction() float64 {
	return math.Atan(v.Y() / v.X())
}

func (v Vector2) Norm() Vector { //TODO: TEST
	return v.Div(v.Magnitude())
}

func (v Vector2) Div(n float64) Vector2 { //TODO TEST
	return Vector2{[2]float64{v.X() / n, v.Y() / n}}
}

func (v1 Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{[2]float64{v1.X() + v2.X(), v1.Y() + v2.Y()}}
}

func (v1 Vector2) Sub(v2 Vector2) Vector2 {
	return Vector2{[2]float64{v1.X() - v2.X(), v1.Y() - v2.Y()}}
}

func (v1 Vector2) Dot(v2 Vector2) float64 {
	return (v1.X() * v2.X()) + (v1.Y() * v2.Y())
}

// 3D
func (v Vector3) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X(), 2) + math.Pow(v.Y(), 2) + math.Pow(v.Z(), 2))
}

func (v Vector3) Norm() Vector { //TODO: TEST
	return v.Div(v.Magnitude())
}

func (v Vector3) Div(n float64) Vector3 { //TODO TEST
	return Vector3{[3]float64{v.X() / n, v.Y() / n, v.Z() / n}}
}

func (v1 Vector3) Add(v2 Vector3) Vector3 { //TODO: TEST
	return NewVector3(v1.X()+v2.X(), v1.Y()+v2.Y(), v1.Z()+v2.Z())
}

func (v1 Vector3) Sub(v2 Vector3) Vector3 { //TODO: TEST
	return NewVector3(v1.X()-v2.X(), v1.Y()-v2.Y(), v1.Z()-v2.Z())
}

func (v1 Vector3) Dot(v2 Vector3) float64 {
	return (v1.X() * v2.X()) + (v1.Y() * v2.Y()) + (v1.Z() * v2.Z())
}

func (a Vector3) Cross(b Vector3) Vector3 {
	x := (a.Y() * b.Z()) - (a.Z() * b.Y())
	y := (a.Z() * b.X()) - (a.X() * b.Z())
	z := (a.X() * b.Y()) - (a.Y() * b.X())
	return NewVector3(x, y, z)
}

func radToDeg(r float64) float64 {
	return r * (180 / math.Pi)
}
