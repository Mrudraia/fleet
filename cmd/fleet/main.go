package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	_ "github.com/golang/glog"
	"github.com/mrudraia/fleetcli/cmd/fleet/account"
	"github.com/mrudraia/fleetcli/pkg/arguments"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:           "fleet",
	Long:          "command line tool for Fleet as a Service(FaaS)",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	err := flag.Set("logtostderr", "true")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't set default error stream: %v\n", err)
		os.Exit(1)
	}

	// register the options that are managed by the flag package
	// pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	fs := root.PersistentFlags()
	// color.AddFlag(root)
	arguments.AddDebugFlag(fs)

	root.AddCommand(account.Cmd)
}

func main() {

	err := flag.CommandLine.Parse([]string{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't parse empty command line to satisfy 'glog': %v\n", err)
		os.Exit(1)
	}
	root.SetArgs(os.Args[1:])
	err = root.Execute()
	if err != nil {
		if !strings.Contains(err.Error(), "Did you mean this?") {
			fmt.Fprintf(os.Stderr, "Failed to execute root command: %s\n", err)
		}
		os.Exit(1)
	}

}
