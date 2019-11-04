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
		cli.BoolFlag{
			Name:  "cross, c",
			Usage: "Cross Product between two 3D vectors",
		},
		cli.BoolFlag{
			Name:  "add, a",
			Usage: "Add two equal-dimension vectors",
		},
		cli.BoolFlag{
			Name:  "sub, s",
			Usage: "Subtract two equal-dimesntion vectors",
		},
		cli.BoolFlag{
			Name:  "norm, n",
			Usage: "Normal of a vector",
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
