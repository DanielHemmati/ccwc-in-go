package main

import (
	"flag"
	"os"
)

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
