/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package self

import (
	"fmt"
	"github.com/stovak/go-terminus/config"

	"github.com/spf13/cobra"
)

func NewSiteVersionCommand(c *config.TerminusConfig) *cobra.Command {
	return &cobra.Command{
		GroupID: "self",
		Use:     "version",
		Short:   "Display the Version Number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(c.GetVersion())
		},
	}
}
