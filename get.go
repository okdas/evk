package main

import (
	"github.com/codegangsta/cli"
	keychain "github.com/okdas/go-keychain"
)

func RunGet(c *cli.Context) error {
	bucket := c.String("bucket")
	environmentVariables, err := getEnvListByBucket(bucket)
	if err != nil {
		return cli.NewExitError("Can't load virtual environments for bucket "+bucket+".\n", 1)
	}

	for i := 0; i < len(environmentVariables); i++ {
		name := environmentVariables[i]
		secret, err := keychain.Find("evk", "evk_"+bucket+"_"+name)
		if err != nil {
			return cli.NewExitError("Can't load secret from a keychain.\n", 1)
		}
		println("export " + name + "=" + secret)
	}

	println("# You can cope and paste those variables to your shell")
	if bucket == "main" {
		println("# or you can run \"eval $(evk get)\"")
	} else {
		println("# or you can run \"eval $(evk get " + bucket + ")\"")
	}
	println("# you also can add it to your .profile to load vars automatically.")

	return nil
}
