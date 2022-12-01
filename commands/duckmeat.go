package commands

import (
	"github.com/npmania/bong/internal/duck"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(duckmeatCmd)
}

var duckmeatCmd = &cobra.Command{
	Use:   "duckmeat",
	Short: "download bangs from duckduckgo",
	Run: func(cmd *cobra.Command, args []string) {
		duck.UpdateBangs()
	},
}
