package lexer

import "github.com/ab36245/go-source/scanner"

type Next[T IToken] func(*scanner.Scanner) T

type Lexer[T IToken] struct {
	scanner *scanner.Scanner
	next    Next[T]
}

func (l *Lexer[T]) Next() T {
	return l.Peek(0)
}

func (l *Lexer[T]) Peek(n int) T {
	return l.next(l.scanner)
}
