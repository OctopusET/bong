package commands

import (
	"fmt"
	"os"

	"github.com/npmania/bong/internal/cli/logsetup"
	"github.com/npmania/bong/internal/tohttps"
	"github.com/spf13/cobra"
)

func init() {
	tohttpsCmd.PersistentFlags().StringVarP(&logsetup.LogLevel, "loglevel", "v", "warn", "Log level (debug, info, warn, error, fatal, panic")
	RootCmd.AddCommand(tohttpsCmd)
}

var tohttpsCmd = &cobra.Command{
	Use:               "tohttps [files]",
	Short:             "Fix http urls to https if supported",
	PersistentPreRunE: logsetup.LoggerSetup,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			tohttps.FilesToHttps(args)
		} else {
			fmt.Println(cmd.UsageString())
			os.Exit(0)
		}
	},
}
