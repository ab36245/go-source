package buffer

import "github.com/ab36245/go-source/input"

func String(str string) Buffer {
	return FromInput(input.String(str))
}
