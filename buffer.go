package source

func newBufferInput(buffer []byte) *bufferInput {
	return &bufferInput{
		buffer: buffer,
		offset: 0,
	}
}

type bufferInput struct {
	buffer []byte
	offset int
}

func (i *bufferInput) Next() (byte, bool) {
	if i.offset < 0 || i.offset >= len(i.buffer) {
		return 0, false
	}
	value := i.buffer[i.offset]
	i.offset++
	return value, true
}
