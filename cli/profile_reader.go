package cli

import (
	"os"
	"strings"
	"user-api/configuration"
)

const (
	pre_fix_profile = "--profile="
)

// ReadProfileArgument return a --profile= setting in application started from line command argument
// Case don't find the --profile= then the default value returned is "production"
func ReadProfileArgument() string {

	args := os.Args
	profile := configuration.Profile_production

	for _, arg := range args {

		if strings.EqualFold(arg, pre_fix_profile+configuration.Profile_local) {
			profile = configuration.Profile_local
		}
	}

	return profile
}
