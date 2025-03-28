package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var printBytes bool

	flag.BoolVar(&printBytes, "c", false, "print the byte/s count")

	flag.Parse()

	if printBytes {
		info, err := Stat("./data/test.txt")

		if err != nil {
			log.Fatalf("could not stat file: %v", err)
		}

		fmt.Printf("File size: %d bytes\n", info.Size())
	}
}
