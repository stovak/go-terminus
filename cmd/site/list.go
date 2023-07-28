package site

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stovak/go-terminus/config"
	"github.com/stovak/go-terminus/pkg/collections"
)

// NewSiteListCommand returns a new cobra command for getting a list of sites
func NewSitesListCommand(tc *config.TerminusConfig) *cobra.Command {
	return &cobra.Command{
		GroupID: "sites",
		Use:     "sites:list",
		Short:   "Get a list of sites",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("sites:list")
			sc := collections.NewSites(tc)
			req := sc.CreateCollectionRequest("GET")
			err := sc.ProcessCollectionResponse(req)
			if err != nil {
				return err
			}
			cmd.Printf("%#v\n", sc)
			return nil
		},
	}
}
