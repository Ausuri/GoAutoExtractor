// File: filewatch/filewatch_tools.go
// Package filewatch provides common useful functions that can be used in interface implementations of filewatch.

package filewatch

import (
	"io/fs"
	"path/filepath"
)

func GetSubDirectories(path string) (directoryPathList []string, err error) {

	var subdirectories []string

	err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && path != "." && path != "/" {
			subdirectories = append(subdirectories, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return subdirectories, nil
}
