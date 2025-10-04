package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Inits a new initilization
func InitViperConfig(cmd *cobra.Command) *viper.Viper {

	// Init a new viper
	viperCfg := viper.New()

	// Set config type. new will be using toml ig
	viperCfg.SetConfigType("toml")

	// Find if passed via cmd flags
	viperCfg.SetConfigFile(cmd.Flags().Lookup("config").Value.String())

	// Read the config Dumbass
	if err := viperCfg.ReadInConfig(); err != nil {
		// Try with env variables, usually when app gets deployed in pods where configs are injected with env variables
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil
		}
	}

	return viperCfg
}
