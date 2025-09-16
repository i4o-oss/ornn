package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConfigurationFile string

type Configuration struct {
	Auth         string
	Cache        string
	Database     string
	FileStorage  string
	Queue        string
	RateLimiting string
	Router       string
	Tasks        bool
	Scheduler    bool
	SMTP         bool
	Webhooks     bool
}

func Init(cmd *cobra.Command) error {
	if ConfigurationFile != "" {
		viper.SetConfigFile(ConfigurationFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.AddConfigPath(home + "/.ornn")
		viper.SetConfigName(".ornn")
		viper.SetConfigType("toml")
	}

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
	}

	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return err
	}

	if viper.ConfigFileUsed() != "" {
		fmt.Printf("Config initialized successfully. Using config file: %s\n", viper.ConfigFileUsed())
	}
	return nil
}
