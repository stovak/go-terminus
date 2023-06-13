package collections

import (
	"github.com/stovak/go-terminus/config"
	"github.com/stovak/go-terminus/pkg/models"
	"net/http"
)

const (
	sitePath = "/api/sites"
)

type Sites struct {
	Collection

	tc    *config.TerminusConfig
	Items map[string]models.Site
}

func NewSites(tc *config.TerminusConfig) *Sites {
	return &Sites{
		tc: tc,
	}
}

func (s *Sites) GetCollectionRequest() *http.Request {
	return s.tc.PrepareRequest("GET", sitePath, nil)
}

func (s *Sites) GetPath() string {
	return sitePath
}
