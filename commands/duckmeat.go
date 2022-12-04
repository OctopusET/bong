package commands

import (
	"github.com/npmania/bong/internal/cli/logsetup"
	"github.com/npmania/bong/internal/thief/duck"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	duckmeatCmd.PersistentFlags().StringVarP(&logsetup.LogLevel, "loglevel", "v", "info", "Log level (debug, info, warn, error, fatal, panic")
	RootCmd.AddCommand(duckmeatCmd)
}

var duckmeatCmd = &cobra.Command{
	Use:               "duckmeat",
	Short:             "Download bangs from duckduckgo",
	PersistentPreRunE: logsetup.LoggerSetup,
	Run: func(cmd *cobra.Command, args []string) {
		if err := duck.UpdateBangs(); err != nil {
			log.WithField("error", err).Error("Failed to update bangs")
		}
	},
}
