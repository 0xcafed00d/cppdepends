package cppdep

import (
	"os"
	"path/filepath"
	"strings"
)

func IsFileType(file string, fileTypes []string) bool {
	for _, v := range fileTypes {
		if strings.HasSuffix(file, v) {
			return true
		}
	}

	return false
}

func ListFiles(rootPath string, fileTypes []string) (files []string, err error) {

	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {

		if err == nil && path != rootPath {
			if info.IsDir() {
				return filepath.SkipDir
			}

			file := info.Name()
			if fileTypes == nil || IsFileType(file, fileTypes) {
				files = append(files, info.Name())
			}
		}
		return nil
	})

	return
}
