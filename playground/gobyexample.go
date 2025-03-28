package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"io"

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

	f, err := os.Open("../data/dani.txt")
	check(err)
	fmt.Println(f)

	defer f.Close()

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, b1)
	// fmt.Printf("%d bytes: %q\n", n1, b1[:n1]) -> show raw bytes
	// fmt.Printf("%d bytes: % x\n", n1, b1[:n1]) -> show hex-dump

	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, io.SeekStart)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()

}
