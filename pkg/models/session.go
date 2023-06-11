package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Session struct {
	Session   string `json:"session"`
	ExpiresAt int64  `json:"expires_at"`
	UserId    string `json:"user_id"`
}

// GetCachedSession returns a session from the cache if it exists
func GetCachedSession() *Session {
	home, _ := os.UserHomeDir()
	contents, _ := os.ReadFile(fmt.Sprintf("%s/.terminus/cache/session", home))
	var toReturn Session
	err := json.Unmarshal(contents, &toReturn)
	if err != nil {
		return &Session{}
	}
	return &toReturn
}

// Validate checks if the session is valid
func (s *Session) Validate() bool {
	if s.Session == "" {
		return false
	}
	// If the expires_at date is in the past, the session is invalid
	if s.ExpiresAt < time.Now().Unix() {
		return false
	}
	return true
}

// AddSessionHeader adds the session headers to the request
func (s *Session) AddSessionHeader(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.Session))
}
