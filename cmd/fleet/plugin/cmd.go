package plugin

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "plugin COMMAND",
	Aliases: []string{"plugins"},
	Short:   "Get information about plugins",
	Long:    "Get information about installed ocm plugins",
	Args:    cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(Cmd)
}
