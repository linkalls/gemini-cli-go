package main

import (
	"fmt"
	"os"

	"github.com/google-gemini/gemini-cli/go-refactor/pkg/cli"
	"github.com/google-gemini/gemini-cli/go-refactor/pkg/config"
)

func main() {
	cfg := config.Parse()

	if cfg.Version {
		fmt.Println("Gemini CLI (Go Refactor) v0.0.1")
		os.Exit(0)
	}

	cli.Run(cfg)
}
