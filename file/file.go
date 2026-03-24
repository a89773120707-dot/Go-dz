package file

import (
	"os"
	"path/filepath"
	"strings"
)

func IsJson(path string) bool {
	ext := filepath.Ext(path)
	ext = strings.ToLower(ext)

	return ext == ".json"

}

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
