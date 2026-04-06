package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	key := os.Getenv("KEY")
	if key == "" {
		return nil, fmt.Errorf("empty KEY in .env")
	}
	return &Config{
		Key: key,
	}, nil
}
