package input

func Buffer(buffer []byte) *Input {
	offset := 0
	return &Input{
		nextByte: func() byte {
			if offset < 0 || offset >= len(buffer) {
				return 0
			}
			value := buffer[offset]
			offset++
			return value
		},
	}
}
