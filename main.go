package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	// "strconv"
	// "strings"
)

func main() {
	var printLines, printWords, printBytes, printChars bool

	flag.BoolVar(&printBytes, "c", false, "print the byte/s count")
	flag.BoolVar(&printLines, "l", false, "print the line count")
	flag.BoolVar(&printWords, "w", false, "print the word count")
	flag.BoolVar(&printChars, "m", false, "print the character count")
	flag.Parse()

	filename := flag.Arg(0)

	var out io.Writer = os.Stdout

	//  if no filename, or there is not stdin then just show the usage
	stat, _ := os.Stdin.Stat()
	if flag.NArg() == 0 && (stat.Mode()&os.ModeCharDevice) != 0 {
		flag.Usage()
		os.Exit(1)
	}

	// if you run `wc ./data/test.txt` you will get only these outputs
	if !printChars && !printWords && !printBytes && !printLines {
		printLines = true
		printWords = true
		printBytes = true
	}

	// file, err := StdinUtil()
	file, err := GetInputReader()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileInfo, err := CountLinesWordsBytes(file)

	if err != nil {
		log.Fatal("error counting lines, words, and bytes")
	}

	if printLines {
		fmt.Fprintf(out, "%d ", fileInfo.line)
	}

	if printWords {
		fmt.Fprintf(out, "%d ", fileInfo.words)
	}

	if printBytes {
		fmt.Fprintf(out, "%d ", fileInfo.bytes)
	}

	if printChars {
		fmt.Fprintf(out, "%d ", fileInfo.chars)
	}

	fmt.Fprintf(out, filename)
}
