package cb

import (
	"encoding/json"

	"github.com/spf13/cobra"
)

type rootOptions struct {
	json bool
}

var rootOpts = &rootOptions{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "cb",
	Long: `cb - The Coinbase v2 CLI client`,
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(
		&rootOpts.json,
		"json",
		"j",
		false,
		"Output data in JSON format instead of the default customized String() format.",
	)
}

func PrettyPrint(data interface{}) ([]byte, error) {
	var p []byte
	
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return p, err
	}

	return p, nil
}
