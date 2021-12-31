package commands

import (
	"github.com/r2r-dev/rules_tar/pkg/size"
	"github.com/spf13/cobra"
)

// SizeCmd ...
var SizeCmd = &cobra.Command{
	Use:   "size",
	Short: "print filezise info",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: verify args
		size.FileSize(args[0])
	},
}
