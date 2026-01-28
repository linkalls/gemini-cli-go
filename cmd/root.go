// Copyright 2026 Google LLC
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"os"

	"github.com/linkalls/gemini-cli-go/pkg/client"
	"github.com/linkalls/gemini-cli-go/pkg/config"
	"github.com/linkalls/gemini-cli-go/pkg/interactive"
	"github.com/spf13/cobra"
)

var (
	prompt       string
	model        string
	outputFormat string
	version      = "0.1.0"
)

var rootCmd = &cobra.Command{
	Use:   "gemini",
	Short: "Gemini CLI - Bring the power of Gemini to your terminal",
	Long: `Gemini CLI is a terminal interface for Google's Gemini AI models.
It provides code understanding, generation, automation, and integration capabilities.`,
	RunE: runGemini,
}

func init() {
	rootCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "Non-interactive mode: provide a prompt")
	rootCmd.Flags().StringVarP(&model, "model", "m", "", "Specify the Gemini model to use")
	rootCmd.Flags().StringVar(&outputFormat, "output-format", "text", "Output format: text, json, or stream-json")
	rootCmd.Version = version
}

func Execute() error {
	return rootCmd.Execute()
}

func runGemini(cmd *cobra.Command, args []string) error {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// Override model if specified via flag
	if model != "" {
		cfg.Model = model
	}

	// Create Gemini client
	geminiClient, err := client.NewClient(cfg)
	if err != nil {
		return fmt.Errorf("failed to create Gemini client: %w", err)
	}
	defer func() {
		if closeErr := geminiClient.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to close client: %v\n", closeErr)
		}
	}()

	// Non-interactive mode with -p flag
	if prompt != "" {
		return runNonInteractive(geminiClient, prompt, outputFormat)
	}

	// Interactive mode
	return interactive.Run(geminiClient, cfg)
}

func runNonInteractive(c *client.Client, prompt string, format string) error {
	response, err := c.GenerateContent(prompt)
	if err != nil {
		return fmt.Errorf("failed to generate content: %w", err)
	}

	switch format {
	case "json":
		fmt.Printf("{\"response\": %q}\n", response)
	case "stream-json":
		fmt.Printf("{\"type\": \"response\", \"text\": %q}\n", response)
	default:
		fmt.Println(response)
	}

	return nil
}
