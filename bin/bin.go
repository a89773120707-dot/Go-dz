package bin

import (
	"fmt"
	"time"
)

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func NewBin(id string, private bool, name string) *Bin {
	return &Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

type BinList struct {
	Bins []Bin
}

func NewBinList() *BinList {
	return &BinList{Bins: make([]Bin, 0)}
}

func (b *BinList) Save(bin *Bin) error {
	b.Bins = append(b.Bins, *bin)
	return nil
}

func (b *BinList) Load(id string) (*Bin, error) {
	for i := range b.Bins {
		if id == b.Bins[i].ID {
			return &b.Bins[i], nil
		}
	}
	return nil, fmt.Errorf("Error load Bin")
}
