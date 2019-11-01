package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	vec "vecc/vector"
)

func TestParses2dVectors(t *testing.T) {
	given := []string{"2", "3", "2.1", "0"}
	want := []vec.Vector{vec.NewVector2(2, 3), vec.NewVector2(2.1, 0)}

	got := parseArgs2(given)

	assert.Equal(t, want, got)
}

func TestParses3dVectors(t *testing.T) {
	given := []string{"2", "3", "2.1", "0", "4", "16000000"}
	want := []vec.Vector{vec.NewVector3(2, 3, 2.1), vec.NewVector3(0, 4, 16000000)}

	got := parseArgs3(given)

	assert.Equal(t, want, got)
}
