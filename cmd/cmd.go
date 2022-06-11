package cmd

import (
	"hello/cmd/root"
	_ "hello/cmd/version"

	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
