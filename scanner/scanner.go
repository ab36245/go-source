package scanner

import (
	"fmt"
	"math"
	"strings"

	"github.com/mattn/go-runewidth"

	"github.com/ab36245/go-source/input"
)

type Scanner struct {
	name  string
	input *input.Input
	lines [][]rune
	row   int
	col   int
}

func (s *Scanner) Current() Char {
	return s.Peek(0)
}

func (s *Scanner) Next() Char {
	c := s.Peek(0)
	s.row, s.col = s.ahead(1)
	return c
}

func (s *Scanner) Peek(n int) Char {
	row, col := s.ahead(n)
	c := s.loadChar(row, col)
	return c
}

func (s *Scanner) Show(fromRow, fromCol, toRow, toCol int) string {
	var result string
	if toRow == fromRow {
		code, mark := s.show(fromRow, fromCol, toCol)
		result += code + "\n" + mark + "\n"
	} else {
		code, mark := s.show(fromRow, fromCol, math.MaxInt)
		result += code + "\n" + mark + "\n"
		if toRow > fromRow+1 {
			result += "---\n"
		}
		code, mark = s.show(toRow, -1, toCol)
		result += code + "\n" + mark + "\n"
	}
	return result
}

func (s *Scanner) ahead(n int) (int, int) {
	row := s.row
	col := s.col
	for n > 0 {
		line := s.loadLine(row)
		if line == nil {
			break
		}
		col++
		if col >= len(line) {
			row++
			col = 0
		}
		n--
	}
	return row, col
}

func (s *Scanner) loadChar(row, col int) Char {
	c := Char{
		scanner: s,
		row:     row,
		col:     col,
	}
	line := s.loadLine(row)
	if col < len(line) {
		c.rune = line[col]
	}
	return c
}

func (s *Scanner) loadLine(row int) []rune {
	for row >= len(s.lines) {
		var line []rune
		for {
			r := s.input.Next()
			if r == 0 {
				break
			}
			line = append(line, r)
			if r == '\n' {
				break
			}
		}
		if len(line) == 0 {
			break
		}
		if line[len(line)-1] != '\n' {
			// Add zero rune to signify EOF
			line = append(line, 0)
		}
		s.lines = append(s.lines, line)
	}
	if row >= len(s.lines) {
		return nil
	}
	return s.lines[row]
}

func (s *Scanner) show(row, fromCol, toCol int) (string, string) {
	line := s.loadLine(row)
	if line == nil {
		return "", ""
	}
	// code := ""
	code := fmt.Sprintf("%03d: ", row+1)
	// mark := ""
	mark := strings.Repeat(" ", len(code))
	for i, r := range line {
		var s string
		if r == 0 {
			s = "<eof>"
		} else if r == '\n' {
			s = ""
		} else if r == '\t' {
			s = strings.Repeat(" ", 4-(i%4))
		} else {
			s = string(r)
		}
		code += s

		var w int
		if i < fromCol {
			s = " "
		} else if i == fromCol {
			s = "^"
		} else if i < toCol {
			s = "."
		} else if i == toCol {
			s = "^"
		} else {
			s = ""
		}
		if r == 0 {
			w = 5 // <eof>
		} else if r == '\n' {
			w = 1
		} else if r == '\t' {
			w = 4 - (i % 4)
		} else {
			w = runewidth.RuneWidth(r)
		}
		mark += strings.Repeat(s, w)
	}
	return code, mark
}
