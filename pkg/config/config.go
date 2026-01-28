// Copyright 2026 Google LLC
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// Config holds the configuration for the Gemini CLI
type Config struct {
	APIKey    string
	Model     string
	ProjectID string
	UseVertex bool
}

// Load loads the configuration from environment variables and config files
func Load() (*Config, error) {
	cfg := &Config{
		Model: "gemini-2.0-flash-exp", // Default model
	}

	// Check for API key
	cfg.APIKey = os.Getenv("GEMINI_API_KEY")
	if cfg.APIKey == "" {
		cfg.APIKey = os.Getenv("GOOGLE_API_KEY")
	}

	// Check for Vertex AI configuration
	if os.Getenv("GOOGLE_GENAI_USE_VERTEXAI") == "true" {
		cfg.UseVertex = true
		cfg.ProjectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
	}

	// If no API key and not using Vertex AI, return error
	if cfg.APIKey == "" && !cfg.UseVertex {
		return nil, fmt.Errorf("no API key found. Set GEMINI_API_KEY or GOOGLE_API_KEY environment variable, or configure Vertex AI")
	}

	return cfg, nil
}

// GetConfigDir returns the configuration directory path
func GetConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".gemini"), nil
}
