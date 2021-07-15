package cb

import (
	"context"
	"fmt"

	v2 "github.com/mikejoh/coinbase-go/v2"
	"github.com/spf13/cobra"
)

// timeCmd represents the subcommand for `cb time`
var timeCmd = &cobra.Command{
	Use:           "time",
	Short:         "Gets the Coinbase API server time.",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := v2.NewConfig()
		client := v2.NewClient(config)

		ctx := context.TODO()
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
