package cb

import (
	"context"
	"fmt"

	v2 "github.com/mikejoh/coinbase-go/v2"
	"github.com/spf13/cobra"
)

type exchangeRateOptions struct {
	currency string
}

var exchangeRateOpts = &exchangeRateOptions{}

// exchangeRatesCmd represents the subcommand for `cb exchange-rates`
var exchangeRatesCmd = &cobra.Command{
	Use:           "exchange-rates",
	Aliases:       []string{"er"},
	Short:         "Gets the current exchange rates for a given currency.",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := v2.NewConfig()
		client := v2.NewClient(config)

		ctx := context.TODO()
		r, err := client.ExchangeRates(ctx, exchangeRateOpts.currency)
		if err != nil {
			return err
		}

		if rootOpts.json {
			json, err := PrettyPrint(r)
			if err != nil {
				return err
			}

			fmt.Printf("%s \n", json)

			return nil
		}

		fmt.Println(r.String())

		return nil
	},
}

func init() {
	exchangeRatesCmd.PersistentFlags().StringVarP(
		&exchangeRateOpts.currency,
		"currency",
		"c",
		"",
		"The currency to fetch the exchange rates for, example: SEK. Default base currency is USD.",
	)

	rootCmd.AddCommand(exchangeRatesCmd)
}
