package app

import (
	"github.com/inquizarus/golbag/pkg/logging"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func makeRootCmd(log logging.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "golbag",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("doing some info logging")
		},
	}
}
