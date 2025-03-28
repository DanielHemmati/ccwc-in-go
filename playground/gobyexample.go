package main

import (
	// "bufio"
	"fmt"
	// "io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileByGobyExample() {
	// dat, err := os.ReadFile("../data/test.txt")
	// check(err)

	f, err := os.Open("../data/test.txt")
	check(err)
	fmt.Println(f)
	defer f.Close()
}
