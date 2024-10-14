package pg

import "fmt"

// errSQLtoSring - ошибка преобразования запроса
func errSQLtoSring(err error) error {
	return fmt.Errorf("ошибка конвертации sql запроса в строку: %w", err)
}

// errSQLQwery - ошибка выполнения sql запроса
func errSQLQwery(err error) error {
	return fmt.Errorf("ошибка выполнения sql запроса: %w", err)
}

// Ошибка парсинга dsn для подключения к БД
func errDSN(err error) error {
	return fmt.Errorf("ошибка парсинга dsn: %w", err)
}

// Ошибка подключения к базе данных
func errDBConect(err error) error {
	return fmt.Errorf("ошибка подключения к database: %w", err)
}
