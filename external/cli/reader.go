package cli

import (
	"os"
	"strings"
)

const (
	pre_fix_profile    = "--profile="
	profile_local      = "local"
	profile_production = "production"
)

// ReadProfileArgument return a --profile= setting in application started from line command argument
// Case don't find the --profile= then the default value returned is "production"
func ReadProfileArgument() string {

	args := os.Args
	profile := profile_production

	for _, arg := range args {

		if strings.EqualFold(arg, pre_fix_profile+profile_local) {
			profile = profile_local
		}
	}

	return profile
}
