package debug

import "github.com/spf13/pflag"

func AddFlag(flags *pflag.FlagSet) {
	flags.BoolVar(
		&enabled,
		"degug",
		false,
		"Enable debug mode.",
	)
}

var enabled bool
