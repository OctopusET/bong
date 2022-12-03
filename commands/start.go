package commands

import (
	"github.com/npmania/bong/internal/server"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start http server",
	Run: func(cmd *cobra.Command, args []string) {
		httpServer := server.HttpServer{}
		httpServer.Start()
	},
}
