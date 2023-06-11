package config

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

var (
	version = "4.0.0-alpha"
)

type Config interface {
	Get(key string) any
	Set(key string, value any)
	Write(to io.WriteCloser) error
}

func NewConfig() *TerminusConfig {
	return &TerminusConfig{
		cfg:     make(map[string]any),
		Verbose: false,
		Config:  "~/.terminus/config.yaml",
		Version: version,
		Build:   getCommitHash(),
	}
}

type TerminusConfig struct {
	Config  string `mapstructure:"config"`
	Verbose bool   `mapstructure:"verbose"`

	Version string
	Build   string

	cfg map[string]any
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

func getCommitHash() string {
	commitHash, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	if err != nil {
		fmt.Println("Failed to retrieve Git commit hash:", err)
		os.Exit(1)
	}

	return string(commitHash)
}
