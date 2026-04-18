package main

import (
	"3-bin_manager/api"
	"3-bin_manager/bin"
	"3-bin_manager/config"
	"3-bin_manager/file"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Загружаем конфиг (.env)
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Ощибка загрузки конфига: %v", err)
		os.Exit(1)
	}

	// Создаём API клиент
	apiClient := api.NewClient(cfg)

	// Создаём репозиторий и хранилище
	repo := file.NewFileBinRepository(&file.OSFileSystem{}, "newdata")
	store := bin.NewBinList(repo)

	//flag
	create := flag.Bool("create", false, "Создать бин из файла (указать путь к JSON файлу)")
	update := flag.Bool("update", false, "Обновить бин по ID (указать ID)")
	deleteFlag := flag.Bool("delete", false, "Удалить бин по ID (указать ID)")
	get := flag.Bool("get", false, "Получить бин по ID (указать ID)")
	list := flag.Bool("list", false, "Вывести список всех бинов")

	//доп флаги
	filePath := flag.String("filePath", "", "Путь к JSON файлу (для create/update)")
	name := flag.String("name", "", "Имя бина (для create)")
	id := flag.String("id", "", "ID бина (для get/delete/update)")

	//Парсим
	flag.Parse()

	//проверить и вывести дефолтное значние
	if !hasFlag() {
		printHelp()
		os.Exit(1)
	}

	//выбор команд
	switch {

	//list
	case *list:
		if err := store.Load(); err != nil {
			fmt.Printf("Ошибка загрузки: %v\n", err)
			os.Exit(1)
		}
		bins := store.ListBin()
		if len(bins) == 0 {
			fmt.Println("Нет Сохраненных Бинов")
			return
		}
		fmt.Println("Сохранённые бины:")
		for i, b := range bins {
			priv := "публичный"
			if b.Private {
				priv = "приватный"
			}
			fmt.Printf("%d. %-20s | ID: %-30s | %s | %s\n",
				i+1, b.Name, b.ID, priv, b.CreatedAt.Format("2006-01-02 15:04"))
		}

	// create
	case *create:
		if *filePath == "" {
			fmt.Println("Ошибка: укажите --file путь к JSON файлу")
			os.Exit(1)
		}
		if *name == "" {
			fmt.Println("Ошибка: укажите --name имя бина")
			os.Exit(1)
		}

		if !file.IsJson(*filePath) {
			fmt.Println("Ошибка: файл должен иметь расширение .json")
			os.Exit(1)
		}

		fileData, err := os.ReadFile(*filePath)
		if err != nil {
			fmt.Printf("Ошибка чтения файла: %v\n", err)
			os.Exit(1)
		}

		var jsonData map[string]interface{}
		if err := json.Unmarshal(fileData, &jsonData); err != nil {
			fmt.Printf("Ошибка парсинга json: %v", err)
		}

		binID, err := apiClient.CreateBin(jsonData)
		if err != nil {
			fmt.Printf("Ошибка создания бина: %v\n", err)
			os.Exit(1)
		}
		newBin := bin.NewBin(*name, false)
		newBin.ID = binID

		if err := store.AddBin(newBin); err != nil {
			fmt.Printf("Ошибка добавления бина: %v\n", err)
			os.Exit(1)
		}

		if err := store.Save(); err != nil {
			fmt.Printf("Ошибка сохранения: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Бин '%s' создан. ID: %s\n", *name, binID)

	//get
	case *get:
		if *id == "" {
			fmt.Println("Ошибка: укажите --id ID бина")
			os.Exit(1)
		}
		record, err := apiClient.GetBin(*id)
		if err != nil {
			fmt.Printf("Ошибка получения бина: %v\n", err)
			os.Exit(1)
		}
		data, err := json.MarshalIndent(record, "", "	")
		if err != nil {
			fmt.Printf("Ошибка форматирования JSON: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(data))

	//delete
	case *deleteFlag:
		if *id == "" {
			fmt.Println("Ошибка: укажите --id ID бина")
			os.Exit(1)
		}
		if err := store.Load(); err != nil {
			fmt.Printf("Ошибка загрузки: %v\n", err)
			os.Exit(1)
		}
		if err := apiClient.DeleteBin(*id); err != nil {
			fmt.Printf("Ошибка удаления из JsonBin: %v\n", err)
			os.Exit(1)
		}

		if err := store.DeletBin(*id); err != nil {
			fmt.Printf("Ошибка локального удаления: %v\n", err)
			os.Exit(1)
		}
		if err := store.Save(); err != nil {
			fmt.Printf("Ошибка сохранения: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Бин %s удалён\n", *id)

	//update
	case *update:
		if *id == "" {
			fmt.Println("Ошибка: укажите --id ID бина")
			os.Exit(1)
		}
		if *filePath == "" {
			fmt.Println("Ошибка: укажите --file путь к JSON файлу")
			os.Exit(1)
		}

		if !file.IsJson(*filePath) {
			fmt.Println("Ошибка: файл должен иметь расширение .json")
			os.Exit(1)
		}

		fileData, err := os.ReadFile(*filePath)
		if err != nil {
			fmt.Printf("Ошибка чтения файла: %v\n", err)
			os.Exit(1)
		}

		var jsonData map[string]interface{}
		if err := json.Unmarshal(fileData, &jsonData); err != nil {
			fmt.Printf("Ошибка JSON в файле: %v\n", err)
			os.Exit(1)
		}

		binID, err := apiClient.UpdateBin(*id, jsonData)
		if err != nil {
			fmt.Printf("Ошибка обновления бина: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Бин %s обновлен.\n", binID)

	default:
		fmt.Println("Неизвестная команда")
		os.Exit(1)
	}

}

func hasFlag() bool {
	visit := false
	flag.Visit(func(f *flag.Flag) {
		visit = true
	})
	return visit
}

func printHelp() {
	fmt.Println("Использование:")
	fmt.Println("  bin-manager -list                          # Список всех бинов")
	fmt.Println("  bin-manager -create -file f.json -name n   # Создать бин")
	fmt.Println("  bin-manager -update -id 1 -file f.json     # Обновить бин")
	fmt.Println("  bin-manager -delete -id 1                  # Удалить бин")
	fmt.Println("  bin-manager -get -id 1                     # Получить бин")
}
