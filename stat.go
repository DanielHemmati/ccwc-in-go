package main

import (
	"os"
)

func Stat(filePath string) (os.FileInfo, error) {
	info, err := os.Stat(filePath)

	if err != nil {
		return nil, err
	}

	return info, nil
}
