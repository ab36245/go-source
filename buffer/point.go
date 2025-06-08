package buffer

type Point struct {
	Index  int
	Line   Line
	Offset int
}

func (p Point) Show() string {
	return p.Line.Show(p.Offset, p.Offset)
}
