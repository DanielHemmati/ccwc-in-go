package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"unicode"
)

type FileStat struct {
	line  int
	words int
	bytes int
	chars int
}

// i feel like my error handling is not that good
func CountLinesWordsBytes(filePath *os.File) (FileStat, error) {
	var bytes int
	var chars int
	var words int
	var lines int

	reader := bufio.NewReader(filePath)
	inWord := false

	for {
		char, size, err := reader.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("error reading file")
		}

		// accumlates the total size
		bytes += size
		// numbero of chars is differnet from number of bytes
		// if the file contains sth like this: `A😊你`
		// rune count (chars): 3
		// byte count = 1 + 4 + 3 = 8
		// yeah emojis are 4 bytes
		chars++

		if char == '\n' {
			lines++
		}

		// this is the coolest part i have seen until now
		if unicode.IsSpace(char) {
			inWord = false
		} else if !inWord {
			inWord = true
			words++
		}
	}

	return FileStat{line: lines, words: words, chars: chars, bytes: bytes}, nil
}
