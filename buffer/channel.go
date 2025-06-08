package buffer

import "github.com/ab36245/go-source/input"

func Channel(channel <-chan byte) Buffer {
	return FromInput(input.Channel(channel))
}
