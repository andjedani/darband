package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var port string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving for vertex",
	Long:  `No longer description Needed.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve(port)
	},
}

func init() {
	//	stats.ServiceUp()
	serveCmd.Flags().StringVarP(&port, "port", "p", "", "port")
	rootCmd.AddCommand(serveCmd)
}

func serve(port string) {
	log.Info("I'll serve, RelX DUDE")
	log.Info(viper.GetString("port"))
}
