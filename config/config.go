package config

import (
	"context"
	"fmt"
	"io"
	"net/http"
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

func NewConfig(ctx *context.Context) *TerminusConfig {
	return &TerminusConfig{
		ctx:     ctx,
		cfg:     make(map[string]any),
		Verbose: false,
		Config:  "~/.terminus/config.yaml",
		Version: version,
		Timeout: 5 * time.Second,
		Build:   getCommitHash(),
		Session: GetCachedSession(),
	}
}

type TerminusConfig struct {
	Config  string `mapstructure:"config"`
	Verbose bool   `mapstructure:"verbose"`

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

func getCommitHash() string {
	commitHash, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	if err != nil {
		fmt.Println("Failed to retrieve Git commit hash:", err)
		os.Exit(1)
	}

	return string(commitHash)
}

func (tc *TerminusConfig) getClient() *http.Client {
	return &http.Client{
		Timeout: tc.Timeout,
	}
}

func (tc *TerminusConfig) PrepareRequest(m string, u string, b io.Reader) *http.Request {
	req, err := http.NewRequestWithContext(*tc.ctx, m, u, b)
	if err != nil {
		fmt.Println("Failed to create request:", err)
		os.Exit(1)
	}
	tc.Session.AddSessionHeader(req)
	return req
}

func (tc *TerminusConfig) SendRequest(req *http.Request) *http.Response {
	resp, err := tc.getClient().Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		os.Exit(1)
	}
	return resp
}
