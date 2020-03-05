package cmd

import (
	"github.com/kuritka/gsvc/depresolver"
	"github.com/kuritka/gsvc/services/httpsproxy"
	"github.com/kuritka/gsvc/svcrunner"
	"github.com/spf13/cobra"
)

var httpsproxyCmd = &cobra.Command{
	Use:   "httpsproxy",
	Short: "https reverse proxy",
	Long:  `you can use https reverse proxy, when you need access forbidden internet domains in your corporate network`,

	Run: func(cmd *cobra.Command, args []string) {

		proxySettings := depresolver.New().MustResolveHttpsProxy()
		proxy := httpsproxy.New(proxySettings, ctx)
		svcrunner.New(proxy).MustRun()

	},
}

func init() {
	rootCmd.AddCommand(httpsproxyCmd)
}