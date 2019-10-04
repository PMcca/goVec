package main

import (
	"fmt"
	"goVec/vector"
	vec "goVec/vector"
)

func getVec() vector.Vector3 {
	v := vector.Vector3{4, 5, 6}
	println(v.Dot(vec.Vector3{1, 2, 3}))
	return v
}
func main() {
	// v2 := vector.Vector2{2, 3}
	// mag := v2.Magnitude()
	// dir := v2.Direction()
	// dot2 := v2.Dot(vector.Vector2{2.1, 7})

	fmt.Printf("%v", getVec())
}
