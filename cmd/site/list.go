package site

import (
	"github.com/spf13/cobra"
	"github.com/stovak/go-terminus/config"
	"github.com/stovak/go-terminus/pkg/collections"
)

// NewSiteListCommand returns a new cobra command for getting a list of sites
func NewSiteListCommand(tc *config.TerminusConfig) *cobra.Command {
	return &cobra.Command{
		GroupID: "site",
		Use:     "site:list",
		Short:   "Get a list of sites",
		RunE: func(cmd *cobra.Command, args []string) error {
			sc := collections.NewSites(tc)
			return nil
		},
	}
}
