package file

import (
	"3-bin_manager/bin"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileBinRepository struct {
	fs  FileSystem
	dir string
}

func NewFileBinRepository(fs FileSystem, dir string) *FileBinRepository {
	return &FileBinRepository{
		fs:  fs,
		dir: dir,
	}
}
func (f *FileBinRepository) filepath(id string) string {
	return filepath.Join(f.dir, id+".json")
}

func (f *FileBinRepository) Save(b *bin.Bin) error {
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}
	path := f.filepath(b.ID)
	if err := f.fs.Write(path, data); err != nil {
		return err
	}

	return nil
}

func (f *FileBinRepository) Load(id string) (*bin.Bin, error) {
	path := f.filepath(id)
	data, err := f.fs.Read(path)
	if err != nil {
		return nil, err
	}
	var b bin.Bin
	if err := json.Unmarshal(data, &b); err != nil {
		return nil, err
	}
	return &b, nil
}

func (f *FileBinRepository) LoadAll() ([]bin.Bin, error) {
	entries, err := os.ReadDir(f.dir)
	if err != nil {
		if os.IsNotExist(err) {
			if mkErr := os.MkdirAll(f.dir, 0755); mkErr != nil {
				return nil, fmt.Errorf("failed to create dir: %w", mkErr)
			}
			return []bin.Bin{}, nil
		}
		return nil, err
	}

	var bins []bin.Bin

	for _, entry := range entries {
		id := strings.TrimSuffix(entry.Name(), ".json")
		b, err := f.Load(id)
		if err != nil {
			return nil, fmt.Errorf("failed to load %s: %w", id, err)
		}
		bins = append(bins, *b)
	}
	return bins, nil
}
