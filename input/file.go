package input

import (
	"fmt"
	"os"
)

func File(path string) (*Input, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return Reader(file), nil
}
