package source

func StringInput(code string) Input {
	return newBufferInput([]byte(code))
}

func StringScanner(name string, code string) *Scanner {
	return NewScanner(name, StringInput(code))
}
