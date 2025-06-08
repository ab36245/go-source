package buffer

import "math"

type Range struct {
	From Point
	To   Point
}

func (r Range) Show() string {
	if r.From.Line.Number == r.To.Line.Number {
		return r.From.Line.Show(r.From.Offset, r.To.Offset)
	}
	s := r.From.Line.Show(r.From.Offset, math.MaxInt)
	s += r.To.Line.Show(math.MinInt, r.To.Offset)
	return s
}
