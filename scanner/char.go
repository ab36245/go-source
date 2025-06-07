package scanner

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
)

type Char struct {
	rune    rune
	scanner *Scanner
	row     int
	col     int
}

func (c Char) Col() int {
	return c.col
}

func (c Char) Is(r rune) bool {
	return c.rune == r
}

func (c Char) IsAny(rs string) bool {
	return strings.ContainsRune(rs, c.rune)
}

func (c Char) IsDigit() bool {
	return unicode.IsLetter(c.rune)
}

func (c Char) IsEOF() bool {
	return c.rune == 0
}

func (c Char) IsError() bool {
	return c.rune == utf8.RuneError
}

func (c Char) IsLetter() bool {
	return unicode.IsLetter(c.rune)
}

func (c Char) IsPrint() bool {
	return unicode.IsPrint(c.rune)
}

func (c Char) IsSpace() bool {
	return unicode.IsSpace(c.rune)
}

func (c Char) Rune() rune {
	return c.rune
}

func (c Char) Row() int {
	return c.row
}

func (c Char) Scanner() *Scanner {
	return c.scanner
}

func (c Char) Size() int {
	if c.IsEOF() {
		return 0
	}
	if c.IsError() {
		return 1
	}
	return utf8.RuneLen(c.rune)
}

func (c Char) Show() string {
	return c.scanner.Show(c.row, c.col, c.row, c.col)
}

func (c Char) String() string {
	switch {
	case c.IsEOF():
		return "<EOF>"
	case c.IsError():
		return "<ERR>"
	case c.Is('\b'):
		return "<BS>"
	case c.Is('\n'):
		return "<NL>"
	case c.Is('\r'):
		return "<CR>"
	case c.Is('\t'):
		return "<TAB>"
	case c.IsPrint():
		return fmt.Sprintf("'%c'", c.rune)
	default:
		return fmt.Sprintf("0x%02X", c.rune)
	}
}

func (c Char) Width() int {
	if c.IsEOF() {
		return 1
	}
	if c.IsError() {
		return 1
	}
	return runewidth.RuneWidth(c.rune)
}
