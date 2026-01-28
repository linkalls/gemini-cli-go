package config

import (
	"os"

	"github.com/spf13/pflag"
)

type Config struct {
	Prompt  string
	Debug   bool
	Model   string
	Version bool
	APIKey  string
}

func Parse(args []string) (*Config, error) {
	cfg := &Config{}
	fs := pflag.NewFlagSet("gemini", pflag.ContinueOnError)

	fs.StringVarP(&cfg.Prompt, "prompt", "p", "", "Run in non-interactive (headless) mode with the given prompt")
	fs.BoolVarP(&cfg.Debug, "debug", "d", false, "Run in debug mode")
	fs.StringVarP(&cfg.Model, "model", "m", "", "Model to use")
	fs.BoolVarP(&cfg.Version, "version", "v", false, "Print version and exit")
	fs.StringVar(&cfg.APIKey, "api-key", "", "Gemini API Key")

	err := fs.Parse(args)
	if err != nil {
		return nil, err
	}

	// Fallback to env var for API Key
	if cfg.APIKey == "" {
		cfg.APIKey = os.Getenv("GEMINI_API_KEY")
	}

	return cfg, nil
}
