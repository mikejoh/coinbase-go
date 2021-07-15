package cb

import (
	"context"
	"fmt"

	v2 "github.com/mikejoh/coinbase-go/v2"
	"github.com/spf13/cobra"
)

// currenciesCmd represents the subcommand for `cb currencies`
var currenciesCmd = &cobra.Command{
	Use:           "currencies",
	Short:         "Lists the supported currencies.",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := v2.NewConfig()
		client := v2.NewClient(config)

		ctx := context.TODO()
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
