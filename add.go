package main

import (
	"github.com/codegangsta/cli"
	"github.com/mewbak/gopass"
	keychain "github.com/okdas/go-keychain"
)

func RunAdd(c *cli.Context) error {
	if !c.Args().Present() {
		return cli.NewExitError("Environment variable not provided\n", 1)
	}

	name := c.Args().First()
	bucket := c.String("bucket")
	secret, err := gopass.GetPass("Enter secret:\n")
	if err != nil {
		return cli.NewExitError("Error getting secret: "+err.Error(), 1)
	}
	if len(secret) == 0 {
		return cli.NewExitError("Secret can't be blank."+err.Error(), 1)
	}

	err = keychain.Add("evk", "evk_"+bucket+"_"+name, secret)
	if err != nil {
		return cli.NewExitError("Can't write to keychain."+err.Error(), 1)
	}

	println("Added", name)
	return nil
}
