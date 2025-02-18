package cmd

import (
	"github-app/cmd/generate_token"
	"github-app/cmd/version"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "github-app",
		Short: "github-app is a cli tool to interact with github apps",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.github-app/config.yaml)")

	rootCmd.AddCommand(generate_token.GenerateTokenCmd)
	rootCmd.AddCommand(version.VersionCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(strings.Join([]string{home, ".github-app"}, "/"))
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()
}
