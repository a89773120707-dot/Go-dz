package storage

import (
	"3-bin_manager/file"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"time"
)

type BinStorage struct {
	Bin       []byte          `json:"bin"`
	FileName  string          `json:"filename"`
	Timestamp time.Time       `json:"timestamp"`
	Fs        file.FileSystem `json:"-"`
}

func NewBinStorage(filename string, fs file.FileSystem) (*BinStorage, error) {
	data, err := fs.Read(filename)
	if err != nil {
		return nil, fmt.Errorf("Error read file: %v", err)
	}
	binStorage := &BinStorage{
		Bin:       data,
		FileName:  fmt.Sprintf("test_%d.json", rand.IntN(100)),
		Timestamp: time.Now(),
		Fs:        fs,
	}
	return binStorage, nil
}

func (b *BinStorage) Readbin(filename string) error {
	data, err := b.Fs.Read(b.FileName)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, b); err != nil {
		return err
	}
	if err := os.WriteFile("restored_"+b.FileName, b.Bin, 0644); err != nil {
		return err
	}
	return nil
}

func (b *BinStorage) SaveBinStorage(filename string) error {
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return err
	}
	return nil

}

func ReadBinList(filename string, fs file.FileSystem) ([]BinStorage, bool, error) {
	data, err := fs.Read(filename)
	if err != nil {
		return nil, false, err
	}

	var binstorage []BinStorage

	if err := json.Unmarshal(data, &binstorage); err != nil {

		var single BinStorage
		if err2 := json.Unmarshal(data, &single); err2 != nil {
			return nil, false, fmt.Errorf("JSON parse error: %v", err2)
		}

		return []BinStorage{single}, false, nil
	}

	return binstorage, true, nil
}

func Print(list []BinStorage, isArry bool) {
	if len(list) == 0 {

		fmt.Println("Empty")
		return
	}

	if isArry {
		for i, v := range list {
			fmt.Printf("=== Массив (%d элементов) ===\n", len(list))
			fmt.Printf("[%d] %s | %d bytes\n", i+1, v.FileName, len(v.Bin))
		}
	} else {
		v := list[0]
		fmt.Printf("=== Одиночный файл ===\n")
		fmt.Printf("Timestamp: %s\n", v.Timestamp.Format("2006-01-02 15:04:05"))
		fmt.Printf("Bin size: %d bytes\n", len(v.Bin))
		fmt.Printf("Content: %s\n", string(v.Bin))
	}

}


