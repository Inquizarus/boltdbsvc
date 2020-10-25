package app

import (
	"fmt"
	"os"

	"github.com/inquizarus/golbag/pkg/logging"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var port string
var database string

// Execute runs the CLI application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	log := logging.NewLogrusLogger(nil)

	rootCmd = makeRootCmd(log)

	rootCmd.PersistentFlags().StringVarP(&port, portVarName, portVarShortName, portVarDefaultValue, "Define which port should be used.")
	rootCmd.PersistentFlags().StringVarP(&database, databaseVarName, databaseVarShortName, databaseVarDefaultValue, "Define which file to use for database.")

	viper.BindPFlag(portVarName, rootCmd.PersistentFlags().Lookup(portVarName))
	viper.BindPFlag(databaseVarName, rootCmd.PersistentFlags().Lookup(databaseVarName))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			panic(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(configFileName)
	}
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
