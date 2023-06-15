package models

import (
	"strings"

	"github.com/stovak/go-terminus/config"
)

type Environment struct {
	Model
	SiteId          string `json:"site_id"`
	Id              string `json:"id"`
	Name            string `json:"name"`
	Region          string `json:"region"`
	EnvironmentType string `json:"environment_type"`
	IsLocked        bool   `json:"is_locked"`
	IsInitialized   bool   `json:"is_initialized"`
	IsFrozen        bool   `json:"is_frozen"`
	IsTerminated    bool   `json:"is_terminated"`
	IsDevelopment   bool   `json:"is_development"`
	IsMultidev      bool   `json:"is_multidev"`
	IsTest          bool   `json:"is_test"`
}

func NewEnvironment(tc *config.TerminusConfig) *Environment {
	return &Environment{
		Model: Model{
			Path: "sites/{site_id}/environments/{id}",
			Tc:   tc,
		},
	}
}

func (e *Environment) GetPath() string {
	return strings.Replace(strings.Replace(e.Path, "{id}", e.Id, 1), "{site_id}", e.SiteId, 1)
}
