package self

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stovak/go-terminus/config"
)

func NewSelfConfigCommand(c *config.TerminusConfig) *cobra.Command {
	return &cobra.Command{
		GroupID: "self",
		Use:     "config",
		Short:   "Display the configuration file used and it's current values",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("Config file: %s\n", viper.ConfigFileUsed())
			cmd.Printf("Config: %#v\n", c)
		},
	}
}
