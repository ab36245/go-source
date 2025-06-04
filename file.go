package source

import (
	"fmt"
	"os"
)

func FileInput(path string) (Input, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return newReaderInput(file), nil
}

func FileScanner(path string) (*Scanner, error) {
	input, err := FileInput(path)
	if err != nil {
		return nil, err
	}
	return NewScanner(path, input), nil
}
