package commands

import (
	"fmt"

	"github.com/npmania/bong/internal/cli/logsetup"
	"github.com/npmania/bong/internal/thief/duck"
	"github.com/spf13/cobra"
)

func init() {
	duckmeatCmd.PersistentFlags().StringVarP(&logsetup.LogLevel, "loglevel", "v", "warn", "Log level (debug, info, warn, error, fatal, panic")
	RootCmd.AddCommand(duckmeatCmd)
}

var duckmeatCmd = &cobra.Command{
	Use:               "duckmeat",
	Short:             "Download bangs from duckduckgo",
	PersistentPreRunE: logsetup.LoggerSetup,
	Run: func(cmd *cobra.Command, args []string) {
		if err := duck.UpdateBangs(); err != nil {
			fmt.Println("error while updating bangs:", err.Error())
		}
	},
}
