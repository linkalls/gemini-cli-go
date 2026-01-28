// Copyright 2026 Google LLC
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name      string
		envVars   map[string]string
		wantError bool
	}{
		{
			name: "with GEMINI_API_KEY",
			envVars: map[string]string{
				"GEMINI_API_KEY": "test-key",
			},
			wantError: false,
		},
		{
			name: "with GOOGLE_API_KEY",
			envVars: map[string]string{
				"GOOGLE_API_KEY": "test-key",
			},
			wantError: false,
		},
		{
			name:      "without API key",
			envVars:   map[string]string{},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original env vars
			origGemini := os.Getenv("GEMINI_API_KEY")
			origGoogle := os.Getenv("GOOGLE_API_KEY")
			defer func() {
				os.Setenv("GEMINI_API_KEY", origGemini)
				os.Setenv("GOOGLE_API_KEY", origGoogle)
			}()

			// Clear env vars
			os.Unsetenv("GEMINI_API_KEY")
			os.Unsetenv("GOOGLE_API_KEY")

			// Set test env vars
			for k, v := range tt.envVars {
				os.Setenv(k, v)
			}

			cfg, err := Load()

			if tt.wantError {
				if err == nil {
					t.Errorf("Load() expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Load() unexpected error: %v", err)
				}
				if cfg == nil {
					t.Errorf("Load() returned nil config")
				}
				if cfg != nil && cfg.APIKey == "" {
					t.Errorf("Load() returned empty API key")
				}
			}
		})
	}
}

func TestDefaultModel(t *testing.T) {
	os.Setenv("GEMINI_API_KEY", "test-key")
	defer os.Unsetenv("GEMINI_API_KEY")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() failed: %v", err)
	}

	expectedModel := "gemini-2.0-flash-exp"
	if cfg.Model != expectedModel {
		t.Errorf("Load() Model = %v, want %v", cfg.Model, expectedModel)
	}
}
