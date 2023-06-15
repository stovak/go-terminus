package models

import (
	"encoding/json"
	"fmt"
	"github.com/stovak/go-terminus/pkg/request"
	"io"
	"net/http"
	"strings"

	"github.com/stovak/go-terminus/config"
)

type ModelInterface interface {
	GetPath() string
	GetModelRequest() *http.Request
	GetModelResponse(req *http.Request) (*ModelInterface, error)
}

type Model struct {
	Path    string
	Tc      *config.TerminusConfig
	Builder request.Builder
}

func (m *Model) CreateModelRequest(id string) *http.Request {
	req := m.Builder.CreateRequest("GET", m.GetPath(), nil)
	newPath := strings.Replace(req.URL.Path, "{id}", id, 1)
	req.URL.Path = newPath
	return req
}

func (m *Model) ProcessModelResponse(req *http.Request) (*Model, error) {
	resp := m.Builder.SendRequest(req)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error getting site: %s", resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(body, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}

func (m *Model) GetPath() string {
	return ""
}
