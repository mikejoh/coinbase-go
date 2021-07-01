package cb

import (
	"context"
	"fmt"

	v2 "github.com/mikejoh/coinbase-go/v2"
	"github.com/spf13/cobra"
)

// currenciesCmd represents the subcommand for `cb currencies`
var currenciesCmd = &cobra.Command{
	Use:   "currencies",
	Short: "Currencies lists known currencies.",
	RunE: func(cmd *cobra.Command, args []string) error {
		config := v2.NewConfig()

		client := v2.NewClient(config)

		ctx := context.Background()

		c, err := client.Currencies(ctx)
		if err != nil {
			return err
		}

		fmt.Println(c.String())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(currenciesCmd)
}
