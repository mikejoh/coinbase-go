package main

import (
	"context"
	"fmt"

	"github.com/mikejoh/coinbase-go"
	"github.com/spf13/cobra"
)

// timeCmd represents the subcommand for `cb time`
var timeCmd = &cobra.Command{
	Use:           "time",
	Short:         "Gets the Coinbase API server time.",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := coinbase.NewClient(coinbase.NewConfig())

		ctx := context.Background()
		t, err := client.Time(ctx)
		if err != nil {
			return err
		}

		if rootOpts.json {
			json, err := PrettyPrint(t)
			if err != nil {
				return err
			}

			fmt.Printf("%s \n", json)

			return nil
		}

		fmt.Println(t.String())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
}
