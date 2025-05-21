package compression

import "fmt"

type MockDecompressor struct {
	IsError bool
}

func (m *MockDecompressor) Decompress(src, dest string) error {

	if m.IsError {
		return fmt.Errorf("Mock decompression error.")
	}

	return nil
}
