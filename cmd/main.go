package main

import (
	cmd "github.com/r2r-dev/rules_tar/cmd/commands"
	"github.com/r2r-dev/rules_tar/pkg/cli"
)

func main() {
	rootCmd := cmd.RootCmd
	rootCmd.AddCommand(
		cmd.ArchiveCmd,
		cmd.SizeCmd,
	)

	cmd := cli.PrepareBaseCmd(rootCmd)
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
