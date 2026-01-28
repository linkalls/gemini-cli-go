package cli

import (
	"fmt"
	"os"

	"github.com/google-gemini/gemini-cli/go-refactor/pkg/config"
	"github.com/google-gemini/gemini-cli/go-refactor/pkg/gemini"
)

func Run(cfg *config.Config) {
	if cfg.Debug {
		fmt.Printf("Debug mode enabled. Config: %+v\n", cfg)
	}

	client := gemini.NewClient(cfg.Model)

	if cfg.Prompt != "" {
		// Non-interactive mode
		response, err := client.GenerateContent(cfg.Prompt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}
		fmt.Println(response)
	} else {
		// Interactive mode stub
		fmt.Println("Gemini CLI (Go Refactor)")
		fmt.Println("Type your prompt and press Enter. (Ctrl+C to exit)")

		// In a real implementation, we would start a REPL here.
		// For now, just show we entered interactive mode.
	}
}
