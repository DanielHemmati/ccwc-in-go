package main

import (
	"bufio"
	"os"
)

func NumberOfLines(filePath string) (int, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}

	// check to see if we read through the file completely
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}
