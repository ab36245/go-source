package lexer

import "github.com/ab36245/go-source/scanner"

type IToken interface {
	Add(scanner.Char)
	Start(TokenKind)
	Show() string
}

type TokenKind int

const (
	InvalidToken TokenKind = iota
	EOFToken
	ErrorToken

	TokenKinds
)

type Token struct {
	scanner *scanner.Scanner
	fromRow int
	fromCol int
	toRow   int
	toCol   int
	Kind    TokenKind
	Text    string
}

func (t *Token) Add(char scanner.Char) {
	if t.scanner == nil {
		t.scanner = char.Scanner()
		t.fromRow = char.Row()
		t.fromCol = char.Col()
	}
	t.toRow = char.Row()
	t.toCol = char.Col()
	if char.IsEOF() {
		t.Text += "<EOF>"
	} else {
		t.Text += string(char.Rune())
	}
}

func (t *Token) Is(kind TokenKind) bool {
	return t.Kind == kind
}

func (t *Token) IsEOF() bool {
	return t.Is(EOFToken)
}

func (t *Token) IsRune(r rune) bool {
	return t.Text == string(r)
}

func (t *Token) IsString(s string) bool {
	return t.Text == s
}

func (t *Token) Start(kind TokenKind) {
	t.Kind = kind
}

func (t *Token) Show() string {
	return t.scanner.Show(t.fromRow, t.fromCol, t.toRow, t.toCol)
}
