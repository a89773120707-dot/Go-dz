package main

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
	return &BinList{
		Bins: make([]Bin, 0),
	}
}
func main() {
	arr := NewBinList()
	fmt.Println(arr)
}
