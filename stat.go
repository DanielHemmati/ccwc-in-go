package main

import (
	"os"
)

// you can also this
func Stat(filePath string) (os.FileInfo, error) {
	info, err := os.Stat(filePath)

	if err != nil {
		return nil, err
	}

	return info, nil
}
