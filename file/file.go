package file

import (
	"os"
	"path/filepath"
	"strings"
)

type FileSystem interface {
	Read(path string) ([]byte, error)
	Write(path string, data []byte) error
}

type OSFileSystem struct {
}

func (o *OSFileSystem) Read(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (o *OSFileSystem) Write(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func IsJson(path string) bool {
	ext := filepath.Ext(path)
	ext = strings.ToLower(ext)

	return ext == ".json"

}
