package collections

import (
	"github.com/stovak/go-terminus/config"
	"strings"
)

const (
	sitePath = "/api/user/{user_id}/membership/sites"
)

type Sites struct {
	Collection
}

func NewSites(tc *config.TerminusConfig) *Sites {
	return &Sites{
		Collection: Collection{
			tc: tc,
		},
	}
}

func (s *Sites) GetPath() string {
	return strings.Replace(sitePath, "{user_id}", s.tc.Session.UserId, 1)
}
