package input

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
)

type Rune rune

const EOF = Rune(0)
const Error = Rune(utf8.RuneError)

func (r Rune) AsDigit(base int) int {
	switch base {
	case 2:
		if r.IsAny("01") {
			return int(r) - int('0')
		}
	case 8:
		if r.IsAny("01234567") {
			return int(r) - int('0')
		}
	case 10:
		if r.IsAny("0123456789") {
			return int(r) - int('0')
		}
	case 16:
		if r.IsAny("0123456789") {
			return int(r) - int('0')
		}
		if r.IsAny("abcdef") {
			return int(r) - int('a') + 10
		}
		if r.IsAny("ABCDEF") {
			return int(r) - int('A') + 10
		}
	}
	return -1
}

func (r Rune) Is(other rune) bool {
	return rune(r) == other
}

func (r Rune) IsAny(rs string) bool {
	return strings.ContainsRune(rs, rune(r))
}

func (r Rune) IsDigit(base int) bool {
	return r.AsDigit(base) >= 0
}

func (r Rune) IsEOF() bool {
	return r == EOF
}

func (r Rune) IsError() bool {
	return r == Error
}

func (r Rune) IsLetter() bool {
	return unicode.IsLetter(rune(r))
}

func (r Rune) IsLower() bool {
	return unicode.IsLower(rune(r))
}

func (r Rune) IsPrint() bool {
	return unicode.IsPrint(rune(r))
}

func (r Rune) IsSpace() bool {
	return unicode.IsSpace(rune(r))
}

func (r Rune) IsUpper() bool {
	return unicode.IsUpper(rune(r))
}

func (r Rune) String() string {
	switch {
	case r.IsEOF():
		return "<EOF>"
	case r.IsError():
		return "<ERR>"
	case r.Is('\b'):
		return "<BS>"
	case r.Is('\n'):
		return "<NL>"
	case r.Is('\r'):
		return "<CR>"
	case r.Is('\t'):
		return "<TAB>"
	case r.IsPrint():
		return fmt.Sprintf("'%c'", rune(r))
	default:
		return fmt.Sprintf("0x%02X", rune(r))
	}
}

func (r Rune) Width() int {
	if r.IsEOF() {
		return 1
	}
	if r.IsError() {
		return 1
	}
	return runewidth.RuneWidth(rune(r))
}
