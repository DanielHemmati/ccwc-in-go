package main

import (
	"flag"
	"os"
)

// StdinUtil function opens the file provided as a command-line argument
// or returns the standard input if no file is specified.
// there is a better way also: https://stackoverflow.com/questions/22744443/check-if-there-is-something-to-read-on-stdin-in-golang
// but i want with lazy way
func StdinUtil() (*os.File, error) {
	if flag.NArg() == 0 {
		return os.Stdin, nil
	}

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		return nil, err
	}

	return file, nil
}
