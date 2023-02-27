package login

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "login",
	Short: "Log in",
	Long:  "Log in, saving the credentials to the configuration file",
	RunE:  run,
}

func init() {

}

func run(cmd *cobra.Command, argv []string) error {
	return nil
}
