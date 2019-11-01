package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMagnitude2D(t *testing.T) {
	want := 5.0
	got := NewVector2(3, 4).Magnitude()
	assert.Equal(t, want, got, "2D Magnitude is incorrect")
}

func TestDot2D(t *testing.T) {
	want := 11.0
	got := NewVector2(2, 4).Dot(NewVector2(3.5, 1))
	assert.Equal(t, want, got, "2D Dot product is incorrect")
}

func TestDirection2D(t *testing.T) {
	want := 1.1071487177940904
	got := NewVector2(2, 4).Direction()
	assert.Equal(t, want, got, "2D direction is incorrect")
}

func TestNorm2D(t *testing.T) {
	want := 1.0
	got := NewVector2(2.3, 3.4).Norm().Magnitude()
	assert.Equal(t, want, got, "2D norm is incorrect")
}

func TestDiv2D(t *testing.T) {
	want := NewVector2(2, 4)
	got := NewVector2(4, 8).Div(2) // Divide by 2
	assert.Equal(t, want, got, "2D Division is incorrect")
}

func TestAdd2D(t *testing.T) {
	want := NewVector2(3.5, 7)

	v1 := NewVector2(1, 4)
	v2 := NewVector2(2.5, 3)
	got := v1.Add(v2)

	assert.Equal(t, want, got, "2D Addition is incorrect")
}

func TestSubtraction2D(t *testing.T) {
	want := NewVector2(0, 4.2)

	v1 := NewVector2(3, 5)
	v2 := NewVector2(3, 0.8)
	got := v1.Sub(v2)

	assert.Equal(t, want, got, "2D Subtraction is incorrect")
}

func TestDot3D(t *testing.T) {
	want := float64(15)

	v1 := NewVector3(4, 2, 1)
	v2 := NewVector3(2.5, 2, 1)

	got := v1.Dot(v2)
	assert.Equal(t, want, got, "3D Dot product incorrect")
}

func TestCross(t *testing.T) {
	want := NewVector3(4, -5.5, 2)

	v1 := NewVector3(3, 4, 5)
	v2 := NewVector3(1, 2, 3.5)

	got := v1.Cross(v2)
	assert.Equal(t, want, got, "Cross product incorrect")
}
