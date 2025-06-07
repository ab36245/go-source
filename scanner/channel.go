package scanner

import "github.com/ab36245/go-source/input"

func Channel(name string, channel <-chan byte) *Scanner {
	return &Scanner{
		name:  name,
		input: input.Channel(channel),
	}
}
