package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filePath string) (string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content, nil
}

func main() {
	content, err := readFile("./data/test.txt")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(content)
}
