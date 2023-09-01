package cmd

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v8"
	"github.com/spf13/cobra"
	"github.com/ugent-library/people-service/models"
	"github.com/ugent-library/people-service/repository"
	"github.com/ugent-library/people-service/ugentldap"
	"go.uber.org/zap"

	// load .env file if present
	_ "github.com/joho/godotenv/autoload"
)

var (
	config Config
)

var logger *zap.SugaredLogger

var rootCmd = &cobra.Command{
	Use: "people-service",
}

func init() {
	cobra.OnInitialize(initConfig, initLogger)
	cobra.OnFinalize(func() {
		logger.Sync()
	})
}

func initConfig() {
	cobra.CheckErr(env.ParseWithOptions(&config, env.Options{
		Prefix: "PEOPLE_",
	}))
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

func newRepository() (models.Repository, error) {
	return repository.NewRepository(&repository.Config{
		DbUrl:  config.Db.Url,
		AesKey: config.Db.AesKey,
	})
}

func newUgentLdapClient() *ugentldap.Client {
	return ugentldap.NewClient(ugentldap.Config{
		Url:      config.Ldap.Url,
		Username: config.Ldap.Username,
		Password: config.Ldap.Password,
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
