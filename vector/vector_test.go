package vector

import (
	"testing"
)

func TestMagnitude2D(t *testing.T) {
	want := 5.0
	got := Vector2{3, 4}.Magnitude()
	if want != got {
		t.Errorf("Expected {%f} got {%f}", want, got)
	}
}

func TestDot2D(t *testing.T) {
	want := 11.0
	got := Vector2{2, 4}.Dot(Vector2{3.5, 1})
	if want != got {
		t.Errorf("Expected {%f} got {%f}", want, got)
	}
}

func TestDirection2D(t *testing.T) {
	want := 1.1071487177940904}
	got := Vector2{2, 4}.Direction()
	assert(float64(want), float64(got), t)
}

func assert(w, g interface{}, t *testing.T) {
	if w != g {
		t.Errorf("Expected {%v} got {%v}", w, g)
	}
}
