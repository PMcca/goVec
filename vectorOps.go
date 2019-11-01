package main

import (
	"log"
	"strconv"
	vecc "vecc/vector"
)

func dot(v []vecc.Vector, dim int) string {
	var r string
	switch dim {
	case V2D:
		r = fmtFloat(v[0].(vecc.Vector2).Dot(v[1].(vecc.Vector2)))
		break
	case V3D:
		r = fmtFloat(v[0].(vecc.Vector3).Dot(v[1].(vecc.Vector3)))
		break
	default:
		log.Fatal("Error printing dot product.")
	}
	return r
}

func cross(v []vecc.Vector, dim int) vecc.Vector3 {
	if vDim != V3D {
		log.Fatal("Error- No 3D vectors given for Cross Product.")
	}

	return v[0].(vecc.Vector3).Cross(v[1].(vecc.Vector3))
}

func add(v []vecc.Vector, dim int) vecc.Vector {
	var r vecc.Vector
	switch dim {
	case V2D:
		r = v[0].(vecc.Vector2).Add(v[1].(vecc.Vector2))
		break
	case V3D:
		r = v[0].(vecc.Vector3).Add(v[1].(vecc.Vector3))
		break
	default:
		log.Fatal("Error adding vectors.") //TODO: Return error instead?
	}
	return r
}

func fmtFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
