package collections

import (
	"github.com/stovak/go-terminus/pkg/models"
	"strings"

	"github.com/stovak/go-terminus/config"
)

type Sites struct {
	Collection
	Items []models.Site
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
	return strings.Replace(s.Path, "{user_id}", s.Tc.Session.UserId, 1)
}

// Get returns a list of sites from the Terminus API
func (s *Sites) Get() error {
	req := s.CreateCollectionRequest("GET")
	err := s.ProcessCollectionResponse(req)
	if err != nil {
		return err
	}
	return nil
}
