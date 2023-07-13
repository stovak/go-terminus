/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package site

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/stovak/go-terminus/config"
	"github.com/stovak/go-terminus/pkg/models"
)

func NewSiteInfoCommand(c *config.TerminusConfig) *cobra.Command {
	return &cobra.Command{
		GroupID: "site",
		Use:     "site:info <sitename>|<site_id>",
		Short:   "Get basic information for a site",
		Args:    cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE: func(cmd *cobra.Command, args []string) error {
			site_id := args[0]
			if site_id == "" {
				return fmt.Errorf("please provide a site ID")
			}
			var s models.Site
			req := http.Request{
				Method: "GET",
				URL: &url.URL{
					Scheme: "https",
					Host:   "terminus.pantheon.io",
					Path:   fmt.Sprintf("/api/sites/%s", site_id),
				},
				Header: map[string][]string{
					"Content-Type": {"application/json"},
				},
			}
			resp, _ := http.DefaultClient.Do(&req)
			if resp.StatusCode != 200 {
				body := make([]byte, resp.ContentLength)
				resp.Body.Read(body)
				cmd.PrintErrf("Error: %s\n", body)
			}
			defer resp.Body.Close()
			b, err := io.ReadAll(resp.Body)
			// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
			if err != nil {
				log.Fatalln(err)
			}
			err = json.Unmarshal(b, &s)
			if err != nil {
				cmd.PrintErrf("ERROR! Body: %s\n", b)
				cmd.PrintErrf("ERROR! %s\n", err.Error())
				return err
			}
			cmd.Printf("SUCCESS!!! Site: %#v\n Body: %s", s, b)
			return nil
		},
	}
}
