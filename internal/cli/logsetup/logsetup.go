package logsetup

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var LogLevel string

func initLogrus(w io.Writer, lv string) error {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.SetOutput(w)

	l, err := log.ParseLevel(lv)
	if err != nil {
		return err
	}

	log.SetLevel(l)

	log.WithFields(log.Fields{
		"LogLevel": l,
	}).Info("Log level has been set")

	return nil
}

func LoggerSetup(cmd *cobra.Command, args []string) error {
	err := initLogrus(os.Stdout, LogLevel)
	if err != nil {
		return err
	}

	return nil
}
