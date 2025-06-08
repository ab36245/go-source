package buffer

import (
	"github.com/ab36245/go-source/input"
)

type Buffer struct {
	runes  []input.Rune
	starts []int
}

func (b Buffer) Line(index int) Line {
	if index < 0 {
		index = 0
	} else if index > len(b.runes) {
		index = len(b.starts)
	}

	// Find line number by binary search
	// We default to the last line. There is a reason for doing this.
	// Starting `hi` at the second last line allows us to always be able
	// to look at the following line.
	number := len(b.starts) - 1
	{
		lo := 0
		hi := number - 1
		for lo <= hi {
			mid := (lo + hi) / 2
			if index < b.starts[mid] {
				hi = mid - 1
			} else if index >= b.starts[mid+1] {
				lo = mid + 1
			} else {
				number = mid
				break
			}
		}
	}

	var runes []input.Rune
	{
		for i := b.starts[number]; i < len(b.runes); i++ {
			r := b.runes[i]
			runes = append(runes, r)
			if r.Is('\n') {
				break
			}
		}
		if len(runes) > 0 && !runes[len(runes)-1].Is('\n') {
			// Add an EOF if the (last) line is not empty
			runes = append(runes, input.EOF)
		}
	}

	return Line{
		Number: number + 1,
		Runes:  runes,
		Start:  b.starts[number],
		Total:  len(b.starts),
	}
}

func (b Buffer) Point(index int) Point {
	line := b.Line(index)
	offset := index - line.Start

	return Point{
		Index:  index,
		Line:   line,
		Offset: offset,
	}
}

func (b Buffer) Range(from int, to int) Range {
	return Range{
		From: b.Point(from),
		To:   b.Point(to),
	}
}

func (b Buffer) RuneAt(index int) input.Rune {
	if index < 0 || index >= len(b.runes) {
		return input.EOF
	}
	return b.runes[index]
}

func (b Buffer) RuneCount() int {
	return len(b.runes)
}
