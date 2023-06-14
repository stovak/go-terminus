package auth

import (
	"github.com/spf13/cobra"
	"github.com/stovak/go-terminus/config"
)

func NewMachineTokenCommand(tc *config.TerminusConfig) *cobra.Command {
	return &cobra.Command{
		GroupID: "auth",
		Args:    cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Short:   "Create/Update a machine token for terminus",
		Long:    "Create/Update a machine token for terminus",
		RunE: func(cmd *cobra.Command, args []string) error {
			tc.Session.SetMachineToken(args[0])
			return nil
		},
	}
}
