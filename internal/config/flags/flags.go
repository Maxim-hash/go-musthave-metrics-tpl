package flags

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	FlagRunAddr string `env:"ADDRESS"`
}

func ParseFlags() *Config {
	flags := &Config{}
	flag.StringVar(&flags.FlagRunAddr, "a", ":8080", "Server listen address")
	flag.Parse()
	if err := env.Parse(flags); err != nil {
		log.Fatalf("Error parsing environment variables: %v", err)
	}

	return flags
}
