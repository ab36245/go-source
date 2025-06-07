package input

func Channel(channel <-chan byte) *Input {
	return &Input{
		nextByte: func() byte {
			return <-channel
		},
	}
}
