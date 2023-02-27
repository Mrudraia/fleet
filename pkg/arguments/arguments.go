package arguments

import (
	"github.com/mrudraia/fleetcli/pkg/debug"
	"github.com/spf13/pflag"
)

func AddDebugFlag(fs *pflag.FlagSet) {
	debug.AddFlag(fs)
}
