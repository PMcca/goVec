package main

import (
	"github.com/urfave/cli"
	"log"
	"strconv"
	vecc "vecc/vector"
)

func dot(v []vecc.Vector, dim int) (string, error) {
	if len(v) < 2 {
		return "", cli.NewExitError("Error- One vector given for dot product- Requires 2.", 1) //TODO: TEST!!!!!
	}
	var r string
	switch dim {
	case V2D:
		r = fmtFloat(v[0].(vecc.Vector2).Dot(v[1].(vecc.Vector2)))
		break
	case V3D:
		r = fmtFloat(v[0].(vecc.Vector3).Dot(v[1].(vecc.Vector3)))
		break
	default:
		return "", cli.NewExitError("Error printing dot product.", 1)
	}
	return r, nil
}

func cross(v []vecc.Vector, dim int) (vecc.Vector3, error) {
	if len(v) < 2 {
		return vecc.EmptyVector3(), cli.NewExitError("Error- One vector given for cross product- Requires 2.", 1) //TODO: TEST!!!!!
	}

	if vDim != V3D {
		log.Fatal("Error- No 3D vectors given for Cross Product.")
	}

	return v[0].(vecc.Vector3).Cross(v[1].(vecc.Vector3)), nil
}

func add(v []vecc.Vector, dim int) (vecc.Vector, error) {
	if len(v) < 2 {
		return vecc.EmptyVector3(), cli.NewExitError("Error- One vector given for add product- Requires 2.", 1) //TODO: TEST!!!!!
	}
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
	return r, nil
}

func sub(v []vecc.Vector, dim int) (vecc.Vector, error) {
	if len(v) < 2 {
		return vecc.EmptyVector3(), cli.NewExitError("Error- One vector given for subtract product- Requires 2.", 1) //TODO: TEST!!!!!
	}
	var r vecc.Vector
	switch dim {
	case V2D:
		r = v[0].(vecc.Vector2).Sub(v[1].(vecc.Vector2))
		break
	case V3D:
		r = v[0].(vecc.Vector3).Sub(v[1].(vecc.Vector3))
		break
	default:
		log.Fatal("Error adding vectors.") //TODO: Return error instead?
	}
	return r, nil
}

func norm(v vecc.Vector, dim int) (vecc.Vector, error) {
	var r vecc.Vector
	switch dim {
	case V2D:
		r = v.(vecc.Vector2).Norm()
		break
	case V3D:
		r = v.(vecc.Vector3).Norm()
		break
	default:
		return vecc.EmptyVector3(), cli.NewExitError("Error with vector dimension in Norm.", 1) //TODO: TEST!!!!!
	}

	return r, nil
}

func mag(v vecc.Vector, dim int) float64 {
	return v.Magnitude()
}

func fmtFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
