package source

type Input interface {
	Next() (byte, bool)
}
