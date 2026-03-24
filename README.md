# Bin Manager

Менеджер для работы с бинарными файлами и их хранения в JSON формате.

## 📦 Структура проекта

```
3-bin_manager/
├── main.go           # Точка входа
├── storage/          # Работа с JSON хранилищем
│   └── storage.go    # SaveBin, ReadBin, ReadBinList
├── file/             # Работа с файлами
│   └── file.go       # IsJson, ReadFile
├── bin/              # Модели данных
│   └── bin.go        # Bin, BinList структуры
└── api/              # API слой (в разработке)
    └── api.go
```

## 🚀 Быстрый старт

### Требования
- Go 1.25.5 или выше

### Установка и запуск

```bash
cd 3-bin_manager

# Сборка
go build .

# Запуск
go run .
```

## 📖 Использование

### Сохранение bin в JSON

```go
import "3-bin_manager/storage"

// Создание bin из файла
bin, err := storage.NewBin("test.bin")
if err != nil {
    return err
}

// Сохранение в JSON
bin.SaveBin("data.json")
```

### Чтение bin из JSON

```go
// Чтение одного bin
err := bin.ReadBin("data.json")
if err != nil {
    return err
}

// Чтение списка bin
bins, err := bin.ReadBinList("data.json")
```

### Работа с файлами

```go
import "3-bin_manager/file"

// Проверка расширения
if file.IsJson("data.json") {
    // это JSON файл
}

// Чтение файла
data, err := file.ReadFile("path/to/file")
```

## 🧪 Тестирование

```bash
# Сборка без ошибок
go build .

# Запуск
go run .
```

## 📝 Формат JSON

```json
{
  "bin": "base64_encoded_data",
  "filename": "test_123.bin",
  "timestamp": "2026-03-24T20:20:54+03:00"
}
```

## 🔧 Конфигурация

| Параметр | Описание |
|----------|----------|
| `go.mod` | Модуль: `3-bin_manager` |
| `Go version` | 1.25.5 |

## 📄 Лицензия

MIT
