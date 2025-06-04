package source

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
)

type Char struct {
	rune rune
	flag charFlag
	span
}

func (c Char) Is(r rune) bool {
	return c.rune == r
}

func (c Char) IsDigit() bool {
	return unicode.IsLetter(c.rune)
}

func (c Char) IsEOF() bool {
	return c.flag&charEOF == charEOF
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

func (c Char) Rune() rune {
	return c.rune
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

type charFlag uint8

const (
	charEOF charFlag = 1 << iota
)
