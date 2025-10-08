package agent

import (
	"flag"
)

type Flags struct {
	FlagServerAddr     string
	FlagPollInterval   int
	FlagReportInterval int
}

func ParseFlags() *Flags {
	flags := &Flags{}
	flag.StringVar(&flags.FlagServerAddr, "a", "localhost:8080", "Server listen address")
	flag.IntVar(&flags.FlagPollInterval, "p", 2, "Polling interval in seconds")
	flag.IntVar(&flags.FlagReportInterval, "r", 10, "Reporting interval in seconds")

	flag.Parse()

	return flags
}
