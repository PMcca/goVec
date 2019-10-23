package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func cliInit() *cli.App {
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

	app.Action = cliActions

	return app
}

func main() {
	app := cliInit()

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
