package bin

type BinRepository interface {
	Save(b *Bin) error
	Load(id string) (*Bin, error)
	LoadAll() ([]Bin, error)
}
