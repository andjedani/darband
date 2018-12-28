package cmd

import (
	"tange/bigv"
	"tange/common"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	//	"github.com/spf13/viper"
)

//var port string

// serveCmd represents the serve command
var initdbCmd = &cobra.Command{
	Use:   "initdb",
	Short: "Initialize database",
	Long:  `No longer description Needed.`,
	Run: func(cmd *cobra.Command, args []string) {
		initdb()
	},
}

func init() {
	rootCmd.AddCommand(initdbCmd)

}

func initdb() {
	log.Info("Initializing Database")
	db := common.Init()
	bigv.AutoMigrate()
	defer db.Close()
}
