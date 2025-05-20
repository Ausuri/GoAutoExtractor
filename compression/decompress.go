package compression

// Feel free to implement your own decompression library with this interface.
type DecompressorInterface interface {
	Decompress(src, dest string) error
}
