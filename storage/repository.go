package storage

import "3-bin_manager/bin"

type BinRepository interface {
	Save(b *bin.Bin) error
	Load(id string) (*bin.Bin, error)
}
