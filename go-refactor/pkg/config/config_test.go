package config

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	// Save original args
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Test case 1
	os.Args = []string{"cmd", "--prompt", "hello", "--debug"}
	cfg := Parse()
	if cfg.Prompt != "hello" {
		t.Errorf("Expected prompt 'hello', got '%s'", cfg.Prompt)
	}
	if !cfg.Debug {
		t.Errorf("Expected debug true, got false")
	}
}
