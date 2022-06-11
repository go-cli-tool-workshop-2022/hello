package version

import (
	"fmt"

	"hello/cmd/root"
	"hello/version"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "Prints version",
	Aliases: []string{"v"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
