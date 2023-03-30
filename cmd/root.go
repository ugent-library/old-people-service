package cmd

import (
	"fmt"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	configFile string
	config     Config
)

var logger *zap.SugaredLogger

var rootCmd = &cobra.Command{
	Use: "people",
}

func init() {
	viper.SetDefault("nats.url", nats.DefaultURL)
	viper.SetDefault("db.url", "postgres://biblio:biblio@localhost:5432/authority?sslmode=disable")
	viper.SetDefault("api.host", "localhost")
	viper.SetDefault("api.port", 3999)
	viper.SetDefault("api_proxy.host", "localhost")
	viper.SetDefault("api_proxy.port", 4001)

	cobra.OnInitialize(initConfig, initLogger)
	cobra.OnFinalize(func() {
		logger.Sync()
	})

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
		cobra.CheckErr(viper.ReadInConfig())
	}
	cobra.CheckErr(viper.Unmarshal(&config))
}

func initLogger() {
	var l *zap.Logger
	var e error
	if config.Production {
		l, e = zap.NewProduction()
	} else {
		l, e = zap.NewDevelopment()
	}
	cobra.CheckErr(e)
	logger = l.Sugar()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
