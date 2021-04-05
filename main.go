package main

import (
	"log"
	"user-api/cli"
	"user-api/configuration"
)

func main() {

	profile := cli.ReadProfileArgument()
	starter, err := configuration.NewAppStarter(profile)

	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	starter.Start()
}
