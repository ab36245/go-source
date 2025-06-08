package buffer

import (
	"io"

	"github.com/ab36245/go-source/input"
)

func Reader(reader io.Reader) Buffer {
	return FromInput(input.Reader(reader))
}
