package decompresser

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Decompress(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fp := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fp, 0755)
			continue
		}
		if err := os.MkdirAll(filepath.Dir(fp), 0755); err != nil {
			return err
		}
		outFile, err := os.Create(fp)
		if err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
