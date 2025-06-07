package scanner

import "github.com/ab36245/go-source/input"

func String(name string, code string) *Scanner {
	return &Scanner{
		name:  name,
		input: input.String(code),
	}
}
