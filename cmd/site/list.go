package site

import (
	"github.com/spf13/cobra"
	"github.com/stovak/go-terminus/config"
)

// NewSiteListCommand returns a new cobra command for getting a list of sites
func NewSiteListCommand(c *config.TerminusConfig) *cobra.Command {
	return &cobra.Command{
		GroupID: "site",
		Use:     "site:list",
		Short:   "Get a list of sites",
		RunE: func(cmd *cobra.Command, args []string) error {
			sl := NewSiteList(c)
			return nil
		},
	}
}
