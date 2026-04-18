package bin

import (
	"crypto/rand"
	"fmt"
	"os"
	"time"
)

type Bin struct {
	ID        string
	Name      string
	Private   bool
	CreatedAt time.Time
}

func NewBin(name string, private bool) *Bin {
	id := generatID()
	return &Bin{
		ID:        id,
		Name:      name,
		Private:   private,
		CreatedAt: time.Now(),
	}
}

func generatID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

type BinList struct {
	Bins []Bin
	Repo BinRepository
}

func NewBinList(repo BinRepository) *BinList {
	return &BinList{
		Bins: []Bin{},
		Repo: repo,
	}
}

func (b *BinList) AddBin(bin *Bin) error {
	if bin.ID == "" {
		return fmt.Errorf("ID cannot be empty")
	}
	for i := range b.Bins {
		if b.Bins[i].ID == bin.ID {
			return fmt.Errorf("bin with ID %s already exists", bin.ID)
		}
	}
	b.Bins = append(b.Bins, *bin)
	return nil
}
func (b *BinList) DeletBin(id string) error {
	for i := range b.Bins {
		if b.Bins[i].ID == id {
			b.Bins = append(b.Bins[:i], b.Bins[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("bin %s not found", id)
}
func (b *BinList) GetBin(id string) (*Bin, error) {
	for i := range b.Bins {
		if id == b.Bins[i].ID {
			return &b.Bins[i], nil
		}
	}
	return nil, fmt.Errorf("bin %s not found", id)
}

func (b *BinList) UpdateBin(id string, newName string, newPrivate *bool) error {
	for i := range b.Bins {
		if b.Bins[i].ID == id {
			if newName != "" {
				b.Bins[i].Name = newName
			}
			if newPrivate != nil {
				b.Bins[i].Private = *newPrivate
			}
			return nil
		}

	}
	return fmt.Errorf("bin with ID %q not found", id)
}

func (b *BinList) ListBin() []Bin {
	return b.Bins
}

func (b *BinList) Save() error {
	for i := range b.Bins {
		if err := b.Repo.Save(&b.Bins[i]); err != nil {
			return fmt.Errorf("failed to save bin %s: %w", b.Bins[i].ID, err)
		}
	}
	return nil
}

func (b *BinList) Load() error {
	bins, err := b.Repo.LoadAll()
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Create NewBinList...")
			return nil
		}
		return fmt.Errorf("failed to load bins: %w", err)
	}
	b.Bins = bins
	return nil
}
