package buffer

import (
	"github.com/ab36245/go-source/input"
)

func File(path string) (Buffer, error) {
	input, err := input.File(path)
	if err != nil {
		return Buffer{}, err
	}
	return FromInput(input), nil
}
