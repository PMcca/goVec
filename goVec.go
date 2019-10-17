package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func cliSetup() *cli.App {
	app := cli.NewApp()

	app.Name = "GoVec"
	app.Usage = "Perform mathematical calculations on given 2D and 3D vectors."
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "dot, d",
			Usage: "Dot Product between two vectors",
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "John"
		arg := c.Args()
		if len(arg) > 0 {
			name = c.Args().Get(0)
		}

		if c.Bool("dot") != false {
			fmt.Println("Hola ", name)
		} else {
			fmt.Println("Hello ", name)
		}

		return nil
	}

	return app
}

func main() {
	app := cliSetup()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

//v := vec.NewVector2(3, 7)
//fmt.Printf("Vector: %v\n Magnitude: %v\n Norm: %v\n", v, v.Magnitude(), v.Norm())
//
//v2 := vec.NewVector2(8, 2)
//fmt.Printf("v Dot product: %v", v.Dot(v2))
