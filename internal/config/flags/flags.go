package flags

import (
	"flag"
)

type Flags struct {
	FlagRunAddr string
}

func ParseFlags() *Flags {
	flags := &Flags{}
	flag.StringVar(&flags.FlagRunAddr, "a", ":8080", "Server listen address")
	flag.Parse()
	return flags
}
