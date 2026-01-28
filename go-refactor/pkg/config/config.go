package config

import (
	"github.com/spf13/pflag"
)

type Config struct {
	Prompt  string
	Debug   bool
	Model   string
	Version bool
}

func Parse() *Config {
	cfg := &Config{}

	pflag.StringVarP(&cfg.Prompt, "prompt", "p", "", "Run in non-interactive (headless) mode with the given prompt")
	pflag.BoolVarP(&cfg.Debug, "debug", "d", false, "Run in debug mode")
	pflag.StringVarP(&cfg.Model, "model", "m", "", "Model to use")
	pflag.BoolVarP(&cfg.Version, "version", "v", false, "Print version and exit")

	pflag.Parse()

	return cfg
}
