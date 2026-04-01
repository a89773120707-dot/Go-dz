package file

import (
	"3-bin_manager/bin"
	"encoding/json"
	"path/filepath"
)

type FileBinRepository struct {
	fs  FileSystem
	dir string
}

func NewFileBinRepository(fs FileSystem, dir string) *FileBinRepository {
	return &FileBinRepository{
		fs:  fs,
		dir:  dir,
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
