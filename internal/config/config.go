package config

import "github.com/joho/godotenv"

// Load - Парсит файл и загружает переменные среды по указному пути
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
