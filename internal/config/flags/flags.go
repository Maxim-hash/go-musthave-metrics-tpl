package flags

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	FlagRunAddr  string `env:"ADDRESS"`
	FlagLogLevel string `env:"LOG_LEVEL"`
}

func ParseFlags() *Config {
	flags := &Config{}
	flag.StringVar(&flags.FlagRunAddr, "a", ":8080", "Server listen address")
	flag.StringVar(&flags.FlagLogLevel, "l", "info", "Logging level")
	flag.Parse()

	env.Parse(flags)

	return flags
}
