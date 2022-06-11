package root

import (
	"fmt"
	"hello/version"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "hello",
	Short: "hello, " + version.Version,
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("Hello World!")
	},
}
