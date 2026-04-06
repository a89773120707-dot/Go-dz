package main

import (
	"3-bin_manager/api"
	"3-bin_manager/bin"
	"3-bin_manager/config"
	"3-bin_manager/file"
	"3-bin_manager/storage"
	"fmt"
)

func main() {
	//файловая система интрефейс
	fs := &file.OSFileSystem{}
	//репозиторий filebinrepo
	repo := file.NewFileBinRepository(fs, "data")
	//новый Bin struct
	b := bin.NewBin("123", true, "Mybin")
	// save bin
	if err := repo.Save(b); err != nil {
		fmt.Println("save error: ", err)
	}
	fmt.Println("saved bin:", b.ID)

	//загрузка bin
	bindata, err := repo.Load("123")
	if err != nil {
		fmt.Println("Load error: ", err)
		return
	}
	fmt.Println("bindata bin:", bindata.ID, bindata.Name, bindata.Private)

	binstorage, err := storage.NewBinStorage("test.bin", fs)
	if err != nil {
		fmt.Println("Error create bin struct:", err)
		return

	}
	isJson := file.IsJson(binstorage.FileName)
	fmt.Println("isJson: ", isJson)
	if isJson {
		if err := binstorage.SaveBinStorage(binstorage.FileName); err != nil {
			fmt.Println("Error save bin:", err)
			return
		}
		if err := binstorage.Readbin(binstorage.FileName); err != nil {
			fmt.Println("Error read bin:", err)
			return
		}

	} else {
		fmt.Println("isJson: ", isJson)
	}
	list, isArry, err := storage.ReadBinList(binstorage.FileName, fs)
	if err != nil {
		fmt.Println("Error readbinlist list: ", err)
		fmt.Println("Filename:", binstorage.FileName)
		return
	}
	storage.Print(list, isArry)

	//Config
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println("Config error: ", err)
		return
	}
	// Api
	api := api.NewApi(cfg)
	_ = api
	fmt.Println(cfg.Key)

}
