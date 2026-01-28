// Copyright 2026 Google LLC
// SPDX-License-Identifier: Apache-2.0

package interactive

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/linkalls/gemini-cli-go/pkg/client"
	"github.com/linkalls/gemini-cli-go/pkg/config"
)

// Run starts the interactive REPL mode
func Run(c *client.Client, cfg *config.Config) error {
	fmt.Println("Gemini CLI - Interactive Mode")
	fmt.Printf("Using model: %s\n", cfg.Model)
	fmt.Println("Type your message and press Enter. Type /quit or /exit to quit.")
	fmt.Println()

	chat := c.StartChat()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		// Handle commands
		if strings.HasPrefix(input, "/") {
			if handleCommand(input) {
				break
			}
			continue
		}

		// Send message to Gemini
		resp, err := chat.SendMessage(genai.Text(input))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}

		// Print response
		if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
			fmt.Println()
			for _, part := range resp.Candidates[0].Content.Parts {
				fmt.Printf("%v\n", part)
			}
			fmt.Println()
		} else {
			fmt.Println("(No response)")
			fmt.Println()
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}

	return nil
}

// handleCommand processes slash commands
// Returns true if the program should exit
func handleCommand(cmd string) bool {
	cmd = strings.ToLower(cmd)

	switch {
	case cmd == "/quit" || cmd == "/exit":
		fmt.Println("Goodbye!")
		return true
	case cmd == "/help":
		printHelp()
		return false
	case cmd == "/clear":
		// Clear screen (simple version)
		fmt.Print("\033[H\033[2J")
		return false
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		fmt.Println("Type /help for available commands")
		return false
	}
}

func printHelp() {
	fmt.Println("\nAvailable commands:")
	fmt.Println("  /help   - Show this help message")
	fmt.Println("  /quit   - Exit the program")
	fmt.Println("  /exit   - Exit the program")
	fmt.Println("  /clear  - Clear the screen")
	fmt.Println()
}
