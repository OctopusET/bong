package commands

import (
	"github.com/npmania/bong/internal/cli/logsetup"
	"github.com/npmania/bong/internal/server"
	"github.com/spf13/cobra"
)

func init() {
	startCmd.PersistentFlags().StringVarP(&logsetup.LogLevel, "loglevel", "v", "warn", "Log level (debug, info, warn, error, fatal, panic")
	RootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:               "start",
	Short:             "Start http server",
	PersistentPreRunE: logsetup.LoggerSetup,
	Run: func(cmd *cobra.Command, args []string) {
		httpServer := server.HttpServer{}
		httpServer.Start()
	},
}
