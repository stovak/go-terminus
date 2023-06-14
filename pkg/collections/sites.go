package collections

import (
	"github.com/stovak/go-terminus/config"
)

const (
	sitePath = "/api/sites"
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
	return sitePath
}
