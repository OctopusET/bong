package commands

import (
	"fmt"

	"github.com/npmania/bong/internal/duck"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(duckmeatCmd)
}

var duckmeatCmd = &cobra.Command{
	Use:   "duckmeat",
	Short: "Download bangs from duckduckgo",
	Run: func(cmd *cobra.Command, args []string) {
		if err := duck.UpdateBangs(); err != nil {
			fmt.Println("error while updating bangs:", err.Error())
		}
	},
}
