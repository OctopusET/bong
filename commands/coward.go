package commands

import (
	"fmt"

	"github.com/npmania/bong/internal/thief/coward"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(cowardCmd)
}

var cowardCmd = &cobra.Command{
	Use:   "coward",
	Short: "Download bangs from Brave Search",
	Run: func(cmd *cobra.Command, args []string) {
		if err := coward.UpdateBangs(); err != nil {
			fmt.Println("error while updating bangs:", err.Error())
		}
	},
}
