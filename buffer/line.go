package buffer

import (
	"fmt"
	"math"
	"strings"

	"github.com/ab36245/go-source/input"
)

type Line struct {
	Number int
	Runes  []input.Rune
	Start  int
	Total  int
}

func (l Line) Show(from, to int) string {
	code := ""
	mark := ""
	for i, r := range l.Runes {
		c := ""
		w := 0
		if r.Is('\n') {
			c = ""
			w = 1
		} else if r.Is('\t') {
			w = 4 - len(code)%4
			c = strings.Repeat(" ", w)
		} else if r.IsEOF() {
			c = "<EOF>"
			w = 5
		} else {
			c = string(r)
			w = r.Width()
		}
		code += c

		m := ""
		if i < from {
			m = " "
		} else if i == from {
			m = "^"
		} else if i < to {
			m = "."
		} else if i == to {
			m = "^"
		}

		mark += strings.Repeat(m, w)
	}
	width := int(math.Trunc(math.Log10(float64(l.Total)) + 1))

	result := fmt.Sprintf("%*d | %s\n", width, l.Number, code)
	result += fmt.Sprintf("%*s | %s\n", width, "", mark)
	return result
}
