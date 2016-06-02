package main

import (
	"fmt"
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
		fmt.Printf("export %s=\"%s\"\n", name, secret)
	}

	command := "eval $(evk get)"
	if bucket != "main" {
		command = "eval $(evk get " + bucket + ")"
	}
	fmt.Printf("# You can cope and paste those variables to your shell or run: \n# %s\n# you also can add it to your .profile to load vars automatically.\n", command)

	return nil
}
