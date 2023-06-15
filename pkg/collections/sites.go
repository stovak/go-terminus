package collections

import (
	"fmt"
	"strings"

	"github.com/stovak/go-terminus/config"
)

type Sites struct {
	Collection
}

func NewSites(tc *config.TerminusConfig) *Sites {
	return &Sites{
		Collection: Collection{
			Path: "/api/user/{user_id}/membership/sites",
			Tc:   tc,
		},
	}
}

func (s *Sites) GetPath() string {
	fmt.Println("instance of Sites")
	return strings.Replace(s.Path, "{user_id}", s.Tc.Session.UserId, 1)
}
