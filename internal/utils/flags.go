package utils

import (
	"flag"
)

func ParseFlags() {
	flag.Bool("log", false, "enable logging to a file")
}

func IsFlagSet(name string) bool {
	flag.Parse()
	flagValue := flag.Lookup(name)
	return flagValue != nil
}
