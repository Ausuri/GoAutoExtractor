package compression

import (
	"context"
	"io"
	"os"

	"github.com/hashicorp/go-extract"
)

type HashigoExtractor struct{}

func (h *HashigoExtractor) Decompress(src, dest string) error {

	ctx := context.Background()
	file, err := os.Open(src)

	if err != nil {
		return err
	}

	reader := io.Reader(file)
	config := extract.NewConfig()

	defer file.Close()

	if err := extract.Unpack(ctx, dest, reader, config); err != nil {
		return err
	}

	return nil
}
