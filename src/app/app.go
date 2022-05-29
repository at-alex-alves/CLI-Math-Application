package app

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// CreateApp creates the application.
func CreateApp() {
	// function stores the name of the function that will be executed.
	var function string

	// x is the value that will be passed to the function.
	var x float64 = 1

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "function",
				Usage:       "Function that will be executed",
				Destination: &function,
				Required:    true,
			},
			&cli.Float64Flag{
				Name:        "x",
				Usage:       "Value that will be passed to the function",
				Destination: &x,
				Required:    true,
			},
		},
		Action: func(c *cli.Context) error {
			executeFunction(function, x)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
