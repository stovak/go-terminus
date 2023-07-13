/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/stovak/go-terminus/cmd/env"

	"github.com/stovak/go-terminus/cmd/self"
	"github.com/stovak/go-terminus/cmd/site"
	"github.com/stovak/go-terminus/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	tc      = config.NewConfig(context.Background())
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terminus",
	Short: "The Command Line interface to Pantheon.io",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	c := initConfig()
	if c != nil {
		fmt.Println(c)
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.terminus/config)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.AddGroup(&cobra.Group{
		ID:    "site",
		Title: "Site commands",
	})
	rootCmd.AddGroup(&cobra.Group{
		ID:    "self",
		Title: "Terminus' innards",
	})
	rootCmd.AddGroup(&cobra.Group{
		ID:    "env",
		Title: "Environment commands",
	})
	rootCmd.AddCommand(site.NewSiteInfoCommand(tc))
	rootCmd.AddCommand(self.NewSelfConfigCommand(tc))
	rootCmd.AddCommand(self.NewSelfVersionCommand(tc))
	rootCmd.AddCommand(site.NewSiteListCommand(tc))
	rootCmd.AddCommand(env.NewEnvInfoCommand(tc))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() *config.TerminusConfig {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	viper.AddConfigPath(path.Join(home, ".terminus"))
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %s  \n", viper.ConfigFileUsed())
		err := viper.Unmarshal(tc)
		if err != nil {
			rootCmd.Println("Unable to decode into struct ", err.Error())
		}
	}
	if tc.Verbose {
		fmt.Println("Verbose output turned on.")
	}
	return tc
}
