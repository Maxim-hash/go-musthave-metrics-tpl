package main

import (
	"flag"
	"fmt"
)

var (
	flagRunAddr string
)

func ParseFlags() {
	flag.StringVar(&flagRunAddr, "a", "localhost:8080", "Server listen address")
	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Printf("unknown arguments: %v\n", flag.Args())
		return
	}
}
