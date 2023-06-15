package env

import (
	"github.com/spf13/cobra"
	"github.com/stovak/go-terminus/config"
)

func NewEnvInfoCommand(tc *config.TerminusConfig) *cobra.Command {
	return &cobra.Command{
		Use:     "env:info",
		GroupID: "env",
		Short:   "Get information for environment",
		Long:    "Get information about for a given environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
