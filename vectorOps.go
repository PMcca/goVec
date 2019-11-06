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
	switch v[0].(type) {
	case vecc.Vector2:
		r = fmtFloat(v[0].(vecc.Vector2).Dot(v[1].(vecc.Vector2)))
		break
	case vecc.Vector3:
		r = fmtFloat(v[0].(vecc.Vector3).Dot(v[1].(vecc.Vector3)))
		break
	default:
		return "", cli.NewExitError("Error printing dot product.", 1)
	}
	return r, nil
}

func cross(v []vecc.Vector, dim int) (vecc.Vector3, error) {
	if len(v) < 2 {
		return vecc.Vector3{}, cli.NewExitError("Error- One vector given for cross product- Requires 2.", 1) //TODO: TEST!!!!!
	}

	switch v[0].(type) {
	case vecc.Vector3:
		return v[0].(vecc.Vector3).Cross(v[1].(vecc.Vector3)), nil

	default:
		return vecc.Vector3{}, cli.NewExitError("Error- Vectors must be 3D for cross product.", 1)
	}
}

func add(v []vecc.Vector, dim int) (vecc.Vector, error) {
	if len(v) < 2 {
		return vecc.Vector3{}, cli.NewExitError("Error- One vector given for add product- Requires 2.", 1) //TODO: TEST!!!!!
	}
	var r vecc.Vector
	switch v[0].(type) {
	case vecc.Vector2:
		r = v[0].(vecc.Vector2).Add(v[1].(vecc.Vector2))
		break
	case vecc.Vector3:
		r = v[0].(vecc.Vector3).Add(v[1].(vecc.Vector3))
		break
	default:
		log.Fatal("Error adding vectors.") //TODO: Return error instead?
	}
	return r, nil
}

func sub(v []vecc.Vector, dim int) (vecc.Vector, error) {
	if len(v) < 2 {
		return vecc.Vector3{}, cli.NewExitError("Error- One vector given for subtract product- Requires 2.", 1) //TODO: TEST!!!!!
	}
	var r vecc.Vector
	switch v[0].(type) {
	case vecc.Vector2:
		r = v[0].(vecc.Vector2).Sub(v[1].(vecc.Vector2))
		break
	case vecc.Vector3:
		r = v[0].(vecc.Vector3).Sub(v[1].(vecc.Vector3))
		break
	default:
		log.Fatal("Error adding vectors.") //TODO: Return error instead?
	}
	return r, nil
}

func norm(v vecc.Vector, dim int) (vecc.Vector, error) {
	var r vecc.Vector
	switch v.(type) {
	case vecc.Vector2:
		r = v.(vecc.Vector2).Norm()
		break
	case vecc.Vector3:
		r = v.(vecc.Vector3).Norm()
		break
	default:
		return vecc.Vector3{}, cli.NewExitError("Error with vector dimension in Norm.", 1) //TODO: TEST!!!!!
	}

	return r, nil
}

func fmtFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
