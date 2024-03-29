package config

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"time"
)

var (
	version = "4.0.0-alpha"
)

type Config interface {
	Get(key string) any
	Set(key string, value any)
	Write(to io.WriteCloser) error
}

// NewConfig Create a default TerminusConfig object.
// It is the base struct for all terminus configurations
func NewConfig(ctx context.Context) *TerminusConfig {
	home, _ := os.UserHomeDir()
	sess, err := GetCachedSession()
	if err != nil {
		fmt.Println("Failed to retrieve session:", err)
		os.Exit(1)
	}
	return &TerminusConfig{
		ctx:     &ctx,
		cfg:     make(map[string]any),
		Verbose: false,
		Config:  home + "/.terminus/config.yml",
		Host:    "terminus.pantheon.io",
		Version: version,
		Timeout: 30 * time.Second,
		Build:   getCommitHash(),
		Session: sess,
	}
}

type TerminusConfig struct {
	Config  string `mapstructure:"config"`
	Verbose bool   `mapstructure:"verbose"`
	Host    string `mapstructure:"host"`

	Version string
	Build   string
	Timeout time.Duration
	Session *Session

	cfg map[string]any
	ctx *context.Context
}

func (tc *TerminusConfig) Get(key string) any {
	return tc.cfg[key]
}

func (tc *TerminusConfig) Set(key string, value any) {
	tc.cfg[key] = value
}

func (tc *TerminusConfig) Write(to io.WriteCloser) error {
	return nil
}

func (tc *TerminusConfig) GetVersion() string {
	return tc.Version + "+" + tc.Build
}

func (tc *TerminusConfig) GetContext() *context.Context {
	return tc.ctx
}

func (tc *TerminusConfig) CreateRequest(m string, p string, v *url.Values) *http.Request {
	toReturn := http.Request{
		Method: m,
		URL:    &url.URL{Scheme: "https", Host: tc.Host, RawPath: p},
		Header: map[string][]string{
			"Accept": {"application/json"},
		},
	}
	_ = tc.Session.AddAuthHeaderToRequest(&toReturn)
	return &toReturn
}

func getCommitHash() string {
	commitHash, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	if err != nil {
		fmt.Println("Failed to retrieve Git commit hash:", err)
		os.Exit(1)
	}

	return string(commitHash)
}
