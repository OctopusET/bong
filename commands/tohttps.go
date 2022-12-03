package commands

import (
	"fmt"
	"os"

	"github.com/npmania/bong/internal/tohttps"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(tohttpsCmd)
}

var tohttpsCmd = &cobra.Command{
	Use:   "tohttps [files]",
	Short: "fix http urls to https if supported",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			tohttps.FilesToHttps(args)
		} else {
			fmt.Println(cmd.UsageString())
			os.Exit(0)
		}
	},
}
