package main

import (
	"goVec/vector"
)

func getVec() vector.Vector {
	v := vector.Vector3{vector.Vector2{3, 2}, 2}
	return v
}
func main() {
	// v2 := vector.Vector2{2, 3}
	// mag := v2.Magnitude()
	// dir := v2.Direction()
	// dot2 := v2.Dot(vector.Vector2{2.1, 7})
	println("V2's values:", "")
}
