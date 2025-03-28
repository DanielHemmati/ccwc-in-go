package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var printBytes, printLines, printChars bool

	flag.BoolVar(&printBytes, "c", false, "print the byte/s count")

	flag.Parse()

	filename := flag.Arg(0)

	if printBytes {
		info, err := Stat(filename)

		if err != nil {
			log.Fatalf("could not stat file: %v", err)
		}

		fmt.Printf("File size: %d bytes\n", info.Size())
	}
}
