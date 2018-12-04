package cmd

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var port string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving for vertex",
	Long:  `No longer description Needed.`,
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()                // fflags is a *flag.FlagSet
		if fflags.Changed("port") == false { // check if the flag "path" is set
			port = viper.GetString("port")
		}
		serve(port)
	},
}

func init() {
	//	stats.ServiceUp()
	serverFlags := serveCmd.Flags()
	serverFlags.StringVarP(&port, "port", "p", "", "port")
	viper.BindPFlags(serverFlags)
	rootCmd.AddCommand(serveCmd)
}

func handler(RespW http.ResponseWriter, ReqR *http.Request) {
	fmt.Fprintf(RespW, "<html><head><title>Watch Out!</title></head>"+
		"<body><h1>Wilkommen %s</h1> Wake up<br>wake up now!</body></html>", ReqR.URL.Path[1:])

}

func serve(port string) {
	log.Info("I'll serve, RelX DUDE")
	portString := ":" + port
	log.Info("On port " + portString)
	http.HandleFunc("/", handler)
	log.Info(http.ListenAndServe(portString, nil))
}
