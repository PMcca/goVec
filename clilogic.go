package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"strconv"
	vecc "vecc/vector"
)

const (
	V2D = iota
	V3D = iota
)

var vDim int

func cliActions(c *cli.Context) error {
	args := c.Args()
	vecs := make([]vecc.Vector, 0)
	var err error = nil

	vecs, err = parseArgs(args)
	if err != nil {
		return err
	}

	var vRes vecc.Vector = nil // Hold results for stringing together operations

	if len(args) == 1 { // Single vector
		vRes = vecs[0]
	}

	if c.Bool("dot") {
		d, err := dot(vecs, vDim)
		if err != nil {
			return err
		}
		fmt.Println(d)
		return nil
	}

	if c.Bool("cross") { //TODO: TEST
		vRes, err = cross(vecs, vDim)
	} else if c.Bool("add") { //TODO: TEST
		vRes, err = add(vecs, vDim)
	} else if c.Bool("sub") { //TODO: TEST
		vRes, err = sub(vecs, vDim)
	}

	if err != nil {
		return err
	}

	if vRes != nil { //vRes should not be null here
		if c.Bool("norm") {
			vRes, err = norm(vRes, vDim)
			if err != nil {
				return err
			}
			if c.Bool("mag") {
				println(fmtFloat(vRes.Magnitude()))
				return nil
			}
		}
		fmt.Print(printVector(vRes))
	} else {
		return cli.NewExitError("Error- no operations given. Type -h or --help for usage.", 1)
	}

	return nil
}

func parseArgs(arg []string) ([]vecc.Vector, error) {
	vecs := make([]vecc.Vector, 2)
	len := len(arg)

	if len == 4 || len == 2 {
		vecs = parseArgs2(arg)
		vDim = V2D

	} else if len == 3 || len == 6 {
		vecs = parseArgs3(arg)
		vDim = V3D

	} else {
		return nil, cli.NewExitError("Wrong number of args- provide one or two 2D/3D vectors", 1)
	}
	return vecs, nil
}

func parseArgs2(arg []string) []vecc.Vector {
	vecs := make([]vecc.Vector, 0)
	for i := 0; i < len(arg); i += 2 {
		x := parseFloat(arg[i])
		y := parseFloat(arg[i+1])
		v := vecc.NewVector2(x, y)
		vecs = append(vecs, v)
	}
	return vecs
}

func parseArgs3(arg []string) []vecc.Vector {
	vecs := make([]vecc.Vector, 0)
	for i := 0; i < len(arg); i += 3 {
		x := parseFloat(arg[i])
		y := parseFloat(arg[1+i])
		z := parseFloat(arg[2+i])
		v := vecc.NewVector3(x, y, z)
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

func printVector(v vecc.Vector) string { //TODO: Format with equal spaced braces
	if vDim == V2D {
		return fmt.Sprintf("[%s]\n[%s]", fmtFloat(v.X()), fmtFloat(v.Y()))
	} else {
		v3 := v.(vecc.Vector3)
		return fmt.Sprintf("[%s]\n|%s|\n[%s]", fmtFloat(v3.X()), fmtFloat(v3.Y()), fmtFloat(v3.Z()))
	}
}
