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
