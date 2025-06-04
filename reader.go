package source

import "io"

func newReaderInput(reader io.Reader) *readerInput {
	return &readerInput{
		reader: reader,
		buffer: make([]byte, 10),
		length: 0,
		offset: 0,
	}
}

type readerInput struct {
	reader io.Reader
	buffer []byte
	length int
	offset int
}

func (i *readerInput) Next() (byte, bool) {
	if i.offset >= i.length {
		n, _ := i.reader.Read(i.buffer)
		if n == 0 {
			return 0, false
		}
		i.length = n
		i.offset = 0
	}
	value := i.buffer[i.offset]
	i.offset++
	return value, true
}
