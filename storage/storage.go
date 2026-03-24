package storage

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Bin struct {
	Bin       []byte    `json:"bin"`
	FileName  string    `json:"filename"`
	Timestamp time.Time `json:"timestamp"`
}

func NewBin(filename string) (*Bin, error) {
	databin, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error read file")
		return nil, err
	}

	bin := &Bin{
		Bin:       databin,
		FileName:  fmt.Sprintf("test_%d.bin", rand.Intn(1000)),
		Timestamp: time.Now(),
	}

	return bin, nil
}

func (b *Bin) ReadBin(filename string) error {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Error read bin")
	}

	if err := json.Unmarshal(databyte, b); err != nil {
		return fmt.Errorf("error convert to json")
	}

	if err := os.WriteFile("restored_"+b.FileName, b.Bin, 0644); err != nil {
		return fmt.Errorf("Error write file")
	}
	fmt.Println("good write ")
	return nil
}
func (b *Bin) ReadBinList(filename string) ([]Bin, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var bins []Bin

	if err := json.Unmarshal(data, &bins); err != nil {
		return nil, err
	}
	return bins, nil
}

func (b *Bin) SaveBin(filename string) {

	file, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Error convert bin to json")
		return
	}

	if err := os.WriteFile(filename, file, 0644); err != nil {
		fmt.Println("Error write file")
		return
	}

	fmt.Println("Good")
}
