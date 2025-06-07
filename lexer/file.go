package lexer

import "github.com/ab36245/go-source/scanner"

func File[T IToken](path string, next Next[T]) (*Lexer[T], error) {
	scanner, err := scanner.File(path)
	if err != nil {
		return nil, err
	}
	return &Lexer[T]{
		scanner: scanner,
		next:    next,
	}, nil
}
