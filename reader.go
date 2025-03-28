package main

import (
	"bufio"
	"io"
	"os"
	"unicode"
)

func CountLinesWordsBytes(filePath string) (lines, words, bytes int, err error) {
	file, err := os.Open(filePath)

	if err != nil {
		// even though i can just write `return` (naked return) i rather to be explicit
		// until i get use to it
		return 0, 0, 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	inWord := false

	for {
		char, size, err := reader.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, 0, 0, err
		}

		bytes += size

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

	return
}
