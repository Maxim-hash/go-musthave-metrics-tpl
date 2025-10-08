package agent

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	FlagServerAddr     string `env:"ADDRESS"`
	FlagPollInterval   int    `env:"POLL_INTERVAL"`
	FlagReportInterval int    `env:"REPORT_INTERVAL"`
}

func ParseFlags() *Config {
	cfg := &Config{}
	flag.StringVar(&cfg.FlagServerAddr, "a", "localhost:8080", "Server listen address")
	flag.IntVar(&cfg.FlagPollInterval, "p", 2, "Polling interval in seconds")
	flag.IntVar(&cfg.FlagReportInterval, "r", 10, "Reporting interval in seconds")

	flag.Parse()

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("Error parsing environment variables: %v", err)
	}

	return cfg
}
