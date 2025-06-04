package source

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
)

func NewScanner(name string, input Input) *Scanner {
	return &Scanner{
		name:  name,
		input: input,
	}
}

type Scanner struct {
	name  string
	input Input
	lines [][]rune
	row   int
	col   int
}

func (s *Scanner) NextChar() Char {
	for s.row >= len(s.lines) && s.NextLine() {
	}

	c := Char{span: s.point()}
	if s.row < len(s.lines) && s.col < len(s.lines[s.row]) {
		c.rune = s.lines[s.row][s.col]
		s.col++
		if s.col >= len(s.lines[s.row]) {
			s.row++
			s.col = 0
		}
	} else {
		c.flag = charEOF
	}
	return c
}

func (s *Scanner) NextLine() bool {
	var line []rune
	for {
		r, ok := s.NextRune()
		if !ok {
			break
		}
		line = append(line, r)
		if r == '\n' {
			break
		}
	}
	if len(line) == 0 {
		return false
	}
	s.lines = append(s.lines, line)
	return true
}

func (s *Scanner) NextRune() (rune, bool) {
	var r int
	var w int
	b, ok := s.input.Next()
	if !ok {
		r = utf8.RuneError
		w = 0
	} else if b&0x80 == 0x00 {
		r = int(b & 0x7F)
		w = 1
	} else if b&0xE0 == 0xC0 {
		r = int(b & 0x3F)
		w = 2
	} else if b&0xF0 == 0xE0 {
		r = int(b & 0x0F)
		w = 3
	} else if b&0xF8 == 0xF0 {
		r = int(b & 0x07)
		w = 4
	} else {
		r = utf8.RuneError
		w = 1
	}
	for i := 1; i < w; i++ {
		b, ok := s.input.Next()
		if !ok {
			r = utf8.RuneError
			w = i
			break
		}
		if b&0xC0 != 0x80 {
			r = utf8.RuneError
			w = i
			break
		}
		r = (r << 6) | int(b&0x3F)
	}
	return rune(r), w > 0
}

func (s *Scanner) Show(fromRow, fromCol, toRow, toCol int) string {
	var show string
	if toRow == fromRow {
		lines := s.show(fromRow, fromCol, toCol)
		show += lines[0] + "\n" + lines[1] + "\n"
	} else if toRow == fromRow+1 {
		lines := s.show(fromRow, fromCol, math.MaxInt)
		show += lines[0] + "\n" + lines[1] + "\n"
		lines = s.show(toRow, -1, toCol)
		show += lines[0] + "\n" + lines[1] + "\n"
	} else {
		lines := s.show(fromRow, fromCol, math.MaxInt)
		show += lines[0] + "\n" + lines[1] + "\n"
		show += "---\n"
		lines = s.show(toRow, -1, toCol)
		show += lines[0] + "\n" + lines[1] + "\n"
	}
	return show
}

func (s *Scanner) show(row, fromCol, toCol int) []string {
	if row >= len(s.lines) {
		return nil
	}
	line := s.lines[row]
	// code := ""
	code := fmt.Sprintf("%03d: ", row+1)
	// show := ""
	show := strings.Repeat(" ", len(code))
	for i, r := range line {
		var s string
		if r == '\n' {
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
		if r == '\n' {
			w = 1
		} else if r == '\t' {
			w = 4 - (i % 4)
		} else {
			w = runewidth.RuneWidth(r)
		}
		show += strings.Repeat(s, w)
	}
	return []string{code, show}
}

func (s *Scanner) point() span {
	return span{
		scanner: s,
		from: pos{
			Row: s.row,
			Col: s.col,
		},
		to: pos{
			Row: s.row,
			Col: s.col,
		},
	}
}
