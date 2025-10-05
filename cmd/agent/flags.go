package main

import (
	"flag"
	"fmt"
)

var (
	flagServerAddr     string
	flagPollInterval   int
	flagReportInterval int
)

func ParseFlags() {
	flag.StringVar(&flagServerAddr, "a", "localhost:8080", "Server listen address")
	flag.IntVar(&flagPollInterval, "p", 2, "Polling interval in seconds")
	flag.IntVar(&flagReportInterval, "r", 10, "Reporting interval in seconds")

	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Printf("unknown arguments: %v\n", flag.Args())
		return
	}
}
