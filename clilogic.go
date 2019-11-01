package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"strconv"
	v "vecc/vector"
)

const (
	V2D = iota
	V3D = iota
)

var vDim int

func cliActions(c *cli.Context) error {
	args := c.Args()
	vecs := make([]v.Vector, 0)
	var err error = nil

	vecs, err = parseArgs(args)
	if err != nil {
		return err
	}

	var vRes v.Vector = nil // Hold results for stringing together operations

	if c.Bool("dot") {
		fmt.Println(dot(vecs, vDim))
		return nil
	}

	if c.Bool("cross") { //TODO: TEST
		vRes = cross(vecs, vDim)
	} else if c.Bool("add") { //TODO: TEST
		vRes = add(vecs, vDim)
	}

	if vRes != nil {
		fmt.Printf("%v", vRes)
	}

	return nil
}

func parseArgs(arg []string) ([]v.Vector, error) {
	vecs := make([]v.Vector, 2)
	len := len(arg)

	if len == 4 || len == 2 {
		vecs = parseArgs2(arg)
		vDim = V2D

	} else if len == 3 || len == 6 {
		vecs = parseArgs3(arg)
		vDim = V3D

	} else {
		return nil, cli.NewExitError("Wrong number of args- provide 4 numbers", 1)
	}
	return vecs, nil
}

func parseArgs2(arg []string) []v.Vector {
	vecs := make([]v.Vector, 0)
	for i := 0; i < len(arg); i += 2 {
		x := parseFloat(arg[i])
		y := parseFloat(arg[i+1])
		v := v.NewVector2(x, y)
		vecs = append(vecs, v)
	}
	return vecs
}

func parseArgs3(arg []string) []v.Vector {
	vecs := make([]v.Vector, 0)
	for i := 0; i < len(arg); i += 3 {
		x := parseFloat(arg[i])
		y := parseFloat(arg[1+i])
		z := parseFloat(arg[2+i])
		v := v.NewVector3(x, y, z)
		vecs = append(vecs, v)
	}
	return vecs
}

func parseFloat(s string) float64 {
	c, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error parsing %s as float", s))
	}
	return c
}
