package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "evk"
	app.Usage = "Secret Environment Variables manager with OS X Keychain backend"
	app.EnableBashCompletion = true
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "Add a new environment variable to bucket (main if not specified).",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "bucket, b",
					Value: "main",
					Usage: "specified bucket",
				},
			},
			Action: RunAdd,
		},
		{
			Name:  "get",
			Usage: "Get environments from a specific bucket.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "bucket, b",
					Value: "main",
					Usage: "specified bucket",
				},
			},
			Action: RunGet,
		},
	}

	app.Run(os.Args)
}
