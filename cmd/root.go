package cmd

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "API Gateway for Vertex",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
		examples and usage of using your application. For example:
		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tange.yaml)")
	viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
	rootCmd.AddCommand(serveCmd)
	//	rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	//	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	//	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	//	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	//	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	//	viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	//	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	//	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	//	viper.SetDefault("license", "apache")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tange")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	log.SetFormatter(&log.TextFormatter{}) //&log.JSONFormatter{})
	logLevel := viper.GetString("log-level")
	if strings.ToUpper(logLevel) == "INFO" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}
	log.Info("🆙")

}
