package configs

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// FBConfig Viper instance for configs
var TangeConfig *viper.Viper

// Load configs from file
func Load(File string, conf *viper.Viper) error {
	conf.SetConfigName(File)
	log.Debug("Trying to load config")
	if err := conf.ReadInConfig(); err != nil {
		log.WithError(err).Warn("Error Loading config file: ", File, ".yml.")
		return err
	}
	log.Println("Using config file:", conf.ConfigFileUsed())
	conf.WatchConfig()
	return nil
}

// Defaults sets the default values and copes with config file, mode 0=empty, mode 1=test
func Defaults() *viper.Viper {
	viperCfg := viper.New()

	viperCfg.AddConfigPath("$GOBIN/tange/configs")
	viperCfg.AddConfigPath("$GOBIN/configs")
	viperCfg.AddConfigPath("configs")

	viperCfg.SetConfigType("yaml")
	//	viperCfg.SetEnvPrefix("tange")
	//	viperCfg.BindEnv("port")
	viperCfg.AutomaticEnv()

	// Setting defaults for this application
	viperCfg.SetDefault("version", "1.0")

	viperCfg.SetDefault("port", "61613")
	viperCfg.SetDefault("timeout", 5*time.Second)
	viperCfg.SetDefault("vertex-port", 3531)
	viperCfg.SetDefault("vertex-server", "127.0.0.1")
	viperCfg.SetDefault("pos-port", 3710)
	viperCfg.SetDefault("pos-server", "127.0.0.1")

	viperCfg.SetDefault("log-level", "Debug")
	viperCfg.AutomaticEnv()

	return viperCfg
}

// Configs tries to load from file, if not successfull loads defaults
func Config(File string) (*viper.Viper, error) {
	v := Defaults()
	if err := Load(File, v); err != nil {
		return v, err
	}
	return v, nil
}
