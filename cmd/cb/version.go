package main

import (
	"fmt"

	"github.com/mikejoh/coinbase-go"
	"github.com/spf13/cobra"
)

// versionCmd represents the subcommand for `cb version`
var versionCmd = &cobra.Command{
	Use:           "version",
	Short:         "Gets the Coinbase client version.",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Version: %s\n", coinbase.GetVersion())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
