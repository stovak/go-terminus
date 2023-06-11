package models

import "github.com/stovak/go-terminus/config"

const (
	sitePath = "/api/site/{id}?get_stats=true"
)

type Site struct {
	tc              *config.TerminusConfig
	Id              string `json:"id"`
	Name            string `json:"name"`
	Created         int64  `json:"created"`
	CreatedByUserId string `json:"created_by_user_id"`
	OrganizationId  string `json:"organization"`
	Label           string `json:"label"`
	Frozen          bool   `json:"frozen"`
	Locked          bool   `json:"locked"`
}

func (s *Site) GetPath() string {
	return sitePath
}
