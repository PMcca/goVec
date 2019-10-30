package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"reflect"
	"strconv"
	v "vecc/vector"
)

func cliActions(c *cli.Context) error {
	args := c.Args()
	var vecs2 []v.Vector2
	var vecs3 []v.Vector3
	vecs := make([]v.Vector, 0)

	if len(args) == 4 { // 2D vectors
		vecs2 = parseArgs2(args)

	} else if len(args) == 6 { // 3D vectors
		vecs3 = parseArgs3(args)

	} else {
		return cli.NewExitError("Wrong number of args- provide 4 numbers", 1)
	}

	if len(vecs2) != 0 {
		for _, n := range vecs2 {
			println(reflect.TypeOf(n).String())
		}
	}

	if len(vecs3) > 0 {
		for _, n := range vecs3 {
			println(reflect.TypeOf(n).String())
		}
	}

	if c.Bool("dot") {
		fmt.Println("Dot product")
	} else {
		fmt.Println("What?")
	}

	return nil
}

func parseArgs(arg []string) []v.Vector {
	vecs := make([]v.Vector, 2)
	if len(arg) == 4 {
		vecs2 := parseArgs2(arg)
		for i, vec := range vecs2 {

		}
	}
	return vecs
}

func parseArgs2(arg []string, vecs *[]v.Vector) {
	for i := 0; i < len(arg); i += 2 {
		x := parseFloat(arg[i])
		y := parseFloat(arg[i+1])
		v := v.NewVector2(x, y)
		vecs = append(vecs, v)
	}

	x = parseFloat(arg[2])
	y = parseFloat(arg[3])
	v2 := v.NewVector2(x, y)
}

func parseArgs3(arg []string) []v.Vector3 {
	x := parseFloat(arg[0])
	y := parseFloat(arg[1])
	z := parseFloat(arg[2])
	v1 := v.NewVector3(x, y, z)

	x = parseFloat(arg[3])
	y = parseFloat(arg[4])
	z = parseFloat(arg[5])
	v2 := v.NewVector3(x, y, z)

	return []v.Vector3{v1, v2}
}

func parseFloat(s string) float64 {
	c, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error parsing %s as float", s))
	}
	return c
}
