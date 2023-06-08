/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/stovak/go-terminus/cmd"
)

var (
	Version    = "4.0.0-alpha"
	CommitHash string
)

func main() {
	cmd.Execute()
}

func init() {
	commitHash, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	if err != nil {
		fmt.Println("Failed to retrieve Git commit hash:", err)
		os.Exit(1)
	}

	CommitHash = string(commitHash)
}
