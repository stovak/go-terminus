package auth

import (
	"net/url"

	"github.com/spf13/cobra"
	"github.com/stovak/go-terminus/config"
)

func NewLoginCommand(tc *config.TerminusConfig) *cobra.Command {
	return &cobra.Command{
		GroupID: "auth",
		Use:     "auth:session",
		Args: 	 cobra.NoArgs,
		Short:   "Create/Update a session for terminus using saved machine token",
		Long:    "Use a saved machine token to create/update a session for terminus from the default machine token",
		RunE: func(cmd *cobra.Command, args []string) error {
,			if tc.Session.Validate() {
				return nil
			}
			// create a REST request to log into terminus
			req := tc.PrepareRequest("POST", "/api/authorize/machine-token", url.Values{})
			// execute the request
			resp, err := tc.SendRequest(&req)
			// add the session header to the request
			config.GetCachedSession().AddSessionHeader(&req)

		},
	}
}
