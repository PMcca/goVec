package main

import (
	"log"
	"strconv"
	v "vecc/vector"
)

func dot(vecs []v.Vector, dim int) string {
	var r string
	switch dim {
	case V2D:
		r = fmtFloat(vecs[0].(v.Vector2).Dot(vecs[1].(v.Vector2)))
		break
	case V3D:
		r = fmtFloat(vecs[0].(v.Vector3).Dot(vecs[1].(v.Vector3)))
		break
	default:
		log.Fatal("Error printing dot product.")
	}
	return r
}

func fmtFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
