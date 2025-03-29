package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	}

	file, err := StdinUtil()
	if err != nil {
		log.Fatal(err)
	}
	if file != os.Stdin {
		defer file.Close()
	}

	fileInfo, err := CountLinesWordsBytes(file)

	if err != nil {
		log.Fatal("error counting lines, words, and bytes")
	}

	if printLines {
		res = append(res, strconv.Itoa(fileInfo.line))
	}

	if printWords {
		res = append(res, strconv.Itoa(fileInfo.words))
	}

	if printChars {
		res = append(res, strconv.Itoa(fileInfo.chars))
	}

	if printBytes {
		res = append(res, strconv.FormatInt(int64(fileInfo.bytes), 10))
	}

	res = append(res, filename)

	fmt.Println(strings.Join(res, " "))
}
