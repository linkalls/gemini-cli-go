package main

import (
	"fmt"
	"os"

	"github.com/linkalls/gemini-cli-go/pkg/cli"
	"github.com/linkalls/gemini-cli-go/pkg/config"
)

func main() {
	cfg, err := config.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing arguments: %v\n", err)
		os.Exit(1)
	}

	if cfg.Version {
		fmt.Println("Gemini CLI (Go Refactor) v0.0.1")
		os.Exit(0)
	}

	cli.Run(cfg)
}
