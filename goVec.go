package main

import (
	"fmt"
	vec "goVec/vector"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	err := cli.NewApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	v := vec.NewVector2(3, 7)
	fmt.Printf("Vector: %v\n Magnitude: %v\n Norm: %v\n", v, v.Magnitude(), v.Norm())

	v2 := vec.NewVector2(8, 2)
	fmt.Printf("v Dot product: %v", v.Dot(v2))
}
