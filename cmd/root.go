//package provides cobra commands
package cmd

import (
	"os"

	"github.com/kuritka/gext/log"
	"github.com/spf13/cobra"
)

var logger = log.Log

var Verbose bool

var rootCmd = &cobra.Command{
	Short: "gsvc depresolver tools",
	Long:  `gsvc is set of tools and microservices used for various situations. i.e. https reverse proxy etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			logger.Error().Msg("No parameters included")
			_ = cmd.Help()
			os.Exit(0)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		logger.Info().Msg("done..")
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
