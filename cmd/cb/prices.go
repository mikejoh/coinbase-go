package cb

import (
	"context"
	"fmt"

	v2 "github.com/mikejoh/coinbase-go/v2"
	"github.com/spf13/cobra"
)

type pricesOptions struct {
	currencyPair string
	priceType    string
}

var pricesOpts = &pricesOptions{}

// pricesCmd represents the subcommand for `cb prices`
var pricesCmd = &cobra.Command{
	Use:           "prices",
	Short:         "Gets the total price to buy one bitcoin or ether.",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := v2.NewConfig()
		client := v2.NewClient(config)

		ctx := context.TODO()
		p, err := client.Prices(ctx, pricesOpts.currencyPair, pricesOpts.priceType)
		if err != nil {
			return err
		}

		if rootOpts.json {
			json, err := PrettyPrint(p)
			if err != nil {
				return err
			}

			fmt.Printf("%s \n", json)

			return nil
		}

		fmt.Println(p.String())

		return nil
	},
}

func init() {
	pricesCmd.PersistentFlags().StringVarP(
		&pricesOpts.currencyPair,
		"pair",
		"p",
		"",
		"The currency pair to query for, example: BTC-USD.",
	)
	pricesCmd.MarkFlagRequired("pair")

	pricesCmd.PersistentFlags().StringVarP(
		&pricesOpts.priceType,
		"type",
		"t",
		"",
		"The price type to query for, valid options are: buy, sell or spot.",
	)
	pricesCmd.MarkFlagRequired("type")

	rootCmd.AddCommand(pricesCmd)
}
