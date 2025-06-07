package input

import "io"

func Reader(reader io.Reader) *Input {
	buffer := make([]byte, 10)
	length := 0
	offset := 0
	return &Input{
		nextByte: func() byte {
			if offset >= length {
				n, _ := reader.Read(buffer)
				if n == 0 {
					return 0
				}
				length = n
				offset = 0
			}
			value := buffer[offset]
			offset++
			return value
		},
	}
}
