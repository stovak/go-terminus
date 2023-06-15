package models

import (
	"strings"

	"github.com/stovak/go-terminus/config"
)

type Site struct {
	Model
	Id              string `json:"id"`
	Name            string `json:"name"`
	Created         int64  `json:"created"`
	CreatedByUserId string `json:"created_by_user_id"`
	OrganizationId  string `json:"organization"`
	Label           string `json:"label"`
	Frozen          bool   `json:"frozen"`
	Locked          bool   `json:"locked"`
}

func NewSite(tc *config.TerminusConfig) *Site {
	return &Site{
		Model: Model{
			Path: "/api/site/{id}?get_stats=true",
			Tc:   tc,
		},
	}
}

func (s *Site) GetPath() string {
	return strings.Replace(s.Path, "{id}", s.Id, 1)
}
