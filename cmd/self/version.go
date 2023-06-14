/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package self

import (
	"github.com/spf13/cobra"
	"github.com/stovak/go-terminus/config"
)

func NewSelfVersionCommand(c *config.TerminusConfig) *cobra.Command {
	return &cobra.Command{
		GroupID: "self",
		Use:     "version",
		Short:   "Display the Version Number",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(c.GetVersion())
		},
	}
}
