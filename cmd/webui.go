package cmd

import (
	"dns/controller"
	"dns/g"
	"dns/httpGin"
	"dns/model"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// webuiCmd represents the webui command
var webuiCmd = &cobra.Command{
	Use:   "webui",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if g.Etcd_url == nil {
			os.Exit(1)
		}
		if len(g.Etcd_url) == 0 {
			os.Exit(1)
		}
		if g.DBKeyPath == "" {
			g.DBKeyPath = "/skydns"
		} else {
			if !strings.HasPrefix(g.DBKeyPath, "/") {
				g.DBKeyPath = "/" + g.DBKeyPath
			}
		}
		//初始化检测etcd链接情况
		model.OninitCheck()
	},
	Run: func(cmd *cobra.Command, args []string) {
		controller.Oninit()
		httpGin.StartHttp()
	},
}

func init() {
	RootCmd.AddCommand(webuiCmd)
	webuiCmd.Flags().StringSliceVar(&g.Etcd_url, "etcdurl", nil, "etcd url not empty")
	webuiCmd.Flags().StringVar(&g.DBKeyPath, "etcdpath", "", "etcd url not empty")
}
