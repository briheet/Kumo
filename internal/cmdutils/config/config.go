package config

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Inits a new initilization
func InitViperConfig(cmd *cobra.Command) *viper.Viper {

	// Init a new viper
	viperCfg := viper.New()

	// Set config type. new will be using toml ig
	viperCfg.SetConfigFile("toml")

	// Find if passed via cmd flags
	viperCfg.SetConfigName(cmd.Flags().Lookup("config").Value.String())

	// Read config Dumbass

	if err := viperCfg.ReadInConfig(); err != nil {

		// Try with env variables, usually when app gets deployed in pods where configs are injected with env variables
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Errorf("couldn't read config file or even environments variables: %w", err)
			return nil
		}

	}

	return viperCfg

}
