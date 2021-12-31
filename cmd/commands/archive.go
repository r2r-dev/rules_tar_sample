package commands

import (
	"github.com/r2r-dev/rules_tar/pkg/archive"
	"github.com/spf13/cobra"
)

func init() {
	registerFlagsArchiveCmd(ArchiveCmd)
}

func registerFlagsArchiveCmd(cmd *cobra.Command) {
	cmd.Flags().String("output", "archive.tar", "output file")
	cmd.Flags().String("prefix", "", "prefix in archive")
}

// ArchiveCmd ...
var ArchiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive files",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: verify args
		output, _ := cmd.Flags().GetString("output")
		prefix, _ := cmd.Flags().GetString("prefix")
		archive.WriteTar(args, output, prefix)
	},
}
