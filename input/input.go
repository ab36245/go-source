package input

import "unicode/utf8"

type Input struct {
	nextByte func() byte
}

func (i *Input) Next() Rune {
	var r int
	var w int
	b := i.nextByte()
	if b == 0 {
		r = 0
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
	for n := 1; n < w; n++ {
		b := i.nextByte()
		if b&0xC0 != 0x80 {
			r = utf8.RuneError
			w = n
			break
		}
		r = (r << 6) | int(b&0x3F)
	}
	return Rune(r)
}
