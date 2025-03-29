package main

import (
	"log"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	t.Run("read fox file", func(t *testing.T) {
		f, err := os.Open("./data/fox.txt")
		if err != nil {
			log.Fatal(err)
		}
		got, _ := CountLinesWordsBytes(f)
		want := FileStat{line: 1, words: 4, chars: 20, bytes: 20}

		if got != want {
			t.Errorf("got %+v want %+v", got, want)
		}
	})

	t.Run("read test.txt file", func(t *testing.T) {
		f, err := os.Open("./data/test.txt")
		if err != nil {
			log.Fatal(err)
		}
		got, _ := CountLinesWordsBytes(f)
		want := FileStat{line: 791, words: 10859, chars: 61031, bytes: 61031}

		if got != want {
			t.Errorf("got %+v want %+v", got, want)
		}
	})
}
