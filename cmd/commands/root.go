package commands

import (
	"github.com/spf13/cobra"
)

// RootCmd is the root command for tarchiver.
var RootCmd = &cobra.Command{
	Use:   "tarchiver",
	Short: "tarchiver",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return nil
	},
}
