package cmd


import (
	"github.com/kuritka/gext/env"
	"github.com/kuritka/gsvc/depresolver"
	"github.com/spf13/cobra"
)

var httpsproxyCmd = &cobra.Command{
	Use:   "httpsproxy",
	Short: "https reverse proxy",
	Long:  `you can use https reverse proxy, when you need access forbidden internet domains in your corporate network`,

	Run: func(cmd *cobra.Command, args []string) {

		settings := depresolver.New().MustResolveHttpsProxy()

		settings.
	},
}

func init() {
	rootCmd.AddCommand(httpsproxyCmd)
}