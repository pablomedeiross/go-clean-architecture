package main

import (
	"log"
	"user-api/external/cli"
	"user-api/external/configuration"

	"github.com/pkg/errors"
)

func main() {

	profile := cli.ReadProfileArgument()
	starter, err := configuration.NewAppStarter(profile)

	if err != nil {
		err := errors.Wrap(err, "Error to start application")
		log.Fatalln(err)
		panic(err)
	}

	starter.Start()
}
