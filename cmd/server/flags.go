package main

import (
	"flag"
)

var (
	flagRunAddr string
)

func ParseFlags() {
	flag.StringVar(&flagRunAddr, "a", "http://localhost:8080", "Server listen address")
	flag.Parse()
}
