package main

import (
	"3-bin_manager/file"
	"3-bin_manager/storage"
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	bin, err := storage.NewBin("test.bin")
	if err != nil {
		fmt.Println("Error create bin struct")
		return
	}
	filename := "data_" + strconv.Itoa(rand.Intn(1000)) + ".json"
	bin.SaveBin(filename)
	bin.ReadBin(filename)

	res := file.IsJson(filename)
	fmt.Println(res)
	if !res {
		fmt.Println("not json")
	}
	file.ReadFile(filename)
	fmt.Println("Read file")
	fmt.Println(filename)
}
