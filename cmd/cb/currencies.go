package main

import (
	"context"
	"fmt"

	"github.com/mikejoh/coinbase-go"
	"github.com/spf13/cobra"
)

// currenciesCmd represents the subcommand for `cb currencies`
var currenciesCmd = &cobra.Command{
	Use:           "currencies",
	Short:         "Lists the supported currencies.",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := coinbase.NewClient(coinbase.NewConfig())

		ctx := context.Background()
		c, err := client.Currencies(ctx)
		if err != nil {
			return err
		}

		if rootOpts.json {
			json, err := PrettyPrint(c)
			if err != nil {
				return err
			}

			fmt.Printf("%s \n", json)

			return nil
		}

		fmt.Println(c.String())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(currenciesCmd)
}
