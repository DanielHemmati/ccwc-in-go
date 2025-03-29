package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var printLines, printWords, printBytes, printChars bool

	flag.BoolVar(&printBytes, "c", false, "print the byte/s count")
	flag.BoolVar(&printLines, "l", false, "print the line count")
	flag.BoolVar(&printWords, "w", false, "print the word count")
	flag.BoolVar(&printChars, "m", false, "print the character count")
	flag.Parse()

	filename := flag.Arg(0)
	var res []string

	if !printChars && !printWords && !printBytes && !printLines {
		printLines = true
		printWords = true
		printBytes = true
		// it's not written in the challenge description
		// but why not
		printChars = true
	}

	fileInfo, err := CountLinesWordsBytes(filename)

	if err != nil {
		log.Fatal("error counting lines, words, and bytes")
	}

	if printBytes {
		stat, err := Stat(filename)
		if err != nil {
			log.Fatal("could not stat file", err)
		}
		res = append(res, strconv.FormatInt(stat.Size(), 10))
	}

	if printChars {
		res = append(res, strconv.Itoa(fileInfo.chars))
	}

	if printWords {
		res = append(res, strconv.Itoa(fileInfo.words))
	}

	if printLines {
		res = append(res, strconv.Itoa(fileInfo.line))
	}

	res = append(res, filename)
	fmt.Println(strings.Join(res, " "))
}
