package source

func newChannelInput(input <-chan byte) *channelInput {
	return &channelInput{
		input: input,
	}
}

type channelInput struct {
	input <-chan byte
}

func (i *channelInput) Next() (byte, bool) {
	value, ok := <-i.input
	return value, ok
}
