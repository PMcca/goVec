package main

import (
	"fmt"
	"github.com/urfave/cli"
)

func cliActions(c *cli.Context) error {
	args := c.Args()
	if len(args) != 2 {
		return cli.NewExitError("Not enough arguments. Provide two vectors of the form {x,y(,z)}", 1)
	}

	if c.Bool("dot") != false {
		fmt.Println("Dot product")
	} else {
		fmt.Println("What?")
	}

	return nil
}
