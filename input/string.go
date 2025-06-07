package input

func String(str string) *Input {
	return Buffer([]byte(str))
}
