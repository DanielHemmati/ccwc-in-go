package main

import (
	"flag"
	"io"
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

// Thanks to: https://x.com/beetbox_games/status/1906145899719372991 suggestion
func GetInputReader() (io.ReadCloser, error) {
	var r io.Reader = os.Stdin
	var err error

	if f := flag.Arg(0); f != "" {
		r, err = os.Open(f)
		if err != nil {
			return nil, err
		}
	}

	return io.NopCloser(r), nil
}
