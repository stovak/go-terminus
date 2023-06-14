package self_test

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stovak/go-terminus/cmd/self"
	"github.com/stovak/go-terminus/config"
)

func TestNewSelfVersionCommand(t *testing.T) {
	tc := config.NewConfig(context.Background())
	cmd := self.NewSelfVersionCommand(tc)
	if cmd.Use != "version" {
		t.Errorf("Expected cmd.Use to be 'version', got '%s'", cmd.Use)
	}
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != tc.GetVersion()+"\n" {
		t.Fatalf("expected \"%s\" got \"%s\"", tc.GetVersion(), string(out))
	}
}
