package utils

import (
	"flag"
)

func ParseFlags() {
	flag.Bool("log", false, "enable logging to a file")

	flag.Parse()
}

func IsFlagSet(name string) bool {
	flagValue := flag.Lookup(name)
	if flagValue == nil {
		return false
	}

	return flagValue.Value.String() == "true"
}
