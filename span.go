package source

type pos struct {
	Row int
	Col int
}

type span struct {
	scanner *Scanner
	from    pos
	to      pos
}

func (s span) FromCol() int {
	return s.from.Col
}

func (s span) FromRow() int {
	return s.from.Row
}

func (s span) ToCol() int {
	return s.to.Col
}

func (s span) ToRow() int {
	return s.to.Row
}

func (s span) Show() string {
	return s.scanner.Show(s.from.Row, s.from.Col, s.to.Row, s.to.Col)
}

func (s span) extend(row, col int) span {
	pos := pos{
		Row: row,
		Col: col,
	}
	return span{
		scanner: s.scanner,
		from:    s.from,
		to:      pos,
	}
}
