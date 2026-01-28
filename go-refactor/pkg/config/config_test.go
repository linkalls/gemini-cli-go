package config

import (
	"testing"
)

func TestParse(t *testing.T) {
	args := []string{"--prompt", "hello", "--debug", "--api-key", "secret"}
	cfg, err := Parse(args)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if cfg.Prompt != "hello" {
		t.Errorf("Expected prompt 'hello', got '%s'", cfg.Prompt)
	}
	if !cfg.Debug {
		t.Errorf("Expected debug true, got false")
	}
	if cfg.APIKey != "secret" {
		t.Errorf("Expected API key 'secret', got '%s'", cfg.APIKey)
	}
}
