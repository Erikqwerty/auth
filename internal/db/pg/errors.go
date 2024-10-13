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
