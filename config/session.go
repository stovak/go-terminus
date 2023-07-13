package config

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"time"
)

const (
	apiPath = "/api/auth/session"
)

type Session struct {
	Session   string `json:"session"`
	ExpiresAt int64  `json:"expires_at"`
	UserId    string `json:"user_id"`
}

type Token struct {
	Token       string `json:"token"`
	Email       string `json:"email"`
	DateCreated int64  `json:"date"`
}

// GetCachedSession returns a session from the cache if it exists
func GetCachedSession() (*Session, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return &Session{}, err
	}
	contents, err := os.ReadFile(path.Join(home, ".terminus", "cache", "session"))
	if err != nil {
		return &Session{}, err
	}
	var toReturn Session
	err = json.Unmarshal(contents, &toReturn)
	if err != nil {
		return &Session{}, err
	}
	fmt.Printf("Session: %#v\n", toReturn)
	return &toReturn, err
}

func GetRemoteSessionFromSavedToken(tc TerminusConfig) (*Session, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return &Session{}, err
	}
	// Get a list of json files in the tokens directory
	files, err := os.ReadDir(path.Join(home, ".terminus", "cache", "tokens"))
	if err != nil {
		return &Session{}, err
	}
	// Let's look through these files for a valid token
	var token Token
	for _, file := range files {
		contents, err := os.ReadFile(path.Join(home, ".terminus", "cache", "tokens", file.Name()))
		// If the directory is unreadable or doesn't exist, error out
		// out and ask for a machine token.
		if err != nil {
			return &Session{}, err
		}
		err = json.Unmarshal(contents, &token)
		if err != nil {
			return &Session{}, err
		}
		// if for some reason this file is empty, skip it
		if token.Empty() {
			continue
		}
		// If the token is not empty, break out of the loop
		break
	}
	// Use this token to build a request for a session
	req := tc.CreateRequest("POST", apiPath, nil)
	req.TransferEncoding = []string{"application/json"}
	return &Session{}, nil
}

// Validate checks if the session is valid
func (s *Session) Validate() bool {
	// If the session is empty, it's invalid
	if s.Session == "" || s.ExpiresAt == 0 || s.UserId == "" {
		return false
	}
	// If the expires_at date is in the past, the session is invalid
	if s.ExpiresAt < time.Now().Unix() {
		return false
	}
	return true
}

// AddSessionHeader adds the session headers to the request
func (s *Session) AddAuthHeaderToRequest(req *http.Request) error {
	if !s.Validate() {
		// if the session is invalid, don't add the header
		return fmt.Errorf("session is invalid")
	}
	req.Header.Add("Authorization", fmt.Sprintf("%s", s.Session))
	return nil
}

func (s *Session) SetMachineToken(s2 string) {

}

func (t *Token) Empty() bool {
	if t.Token == "" {
		return true
	}
	return false
}

func (t *Token) Validate() bool {
	if t.Empty() {
		return false
	}
	if t.DateCreated == 0 {
		return false
	}
	return true
}
