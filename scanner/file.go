package scanner

import "github.com/ab36245/go-source/input"

func File(path string) (*Scanner, error) {
	input, err := input.File(path)
	if err != nil {
		return nil, err
	}
	return &Scanner{
		name:  path,
		input: input,
	}, nil
}
