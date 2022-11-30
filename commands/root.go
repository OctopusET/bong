package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "bong",
	Short: "configurable search engine redirector",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(cmd.UsageString())
			os.Exit(0)
		}
	},
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		if strings.HasPrefix(err.Error(), "unknown command ") {
			os.Exit(1)
		}

		panic(err)
	}
}
