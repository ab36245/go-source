package buffer

import "github.com/ab36245/go-source/input"

func FromInput(in *input.Input) Buffer {
	var lines []int
	var runes []input.Rune

	index := 0
	start := true
	for {
		if start {
			lines = append(lines, index)
		}
		r := in.Next()
		if r.IsEOF() {
			break
		}
		runes = append(runes, r)
		index++
		start = r.Is('\n')
	}
	return Buffer{
		starts: lines,
		runes:  runes,
	}
}
