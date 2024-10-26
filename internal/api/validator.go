package api

import (
	"errors"
	"reflect"

	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
	"github.com/erikqwerty/auth/pkg/utils/validator"
)

// валидируемые данные
const (
	email           = "Email"
	name            = "Name"
	password        = "Password"
	passwordConfirm = "PasswordConfirm"
	role            = "Role"
	id              = "Id"
)

// Ошибки для валидации
var (
	ErrEmailNotSpecified           = errors.New("не указан email")
	ErrNameNotSpecified            = errors.New("имя пользователя указано пустым")
	ErrPasswordNotSpecified        = errors.New("не указан пароль")
	ErrPasswordConfirmNotSpecified = errors.New("не указан пароль подтверждения")
	ErrPasswordsDoNotMatch         = errors.New("пароли не совпадают")
	ErrRoleNotSpecified            = errors.New("роль пользователя не была указана")
	ErrInvalidRole                 = errors.New("переданная роль пользователя не существует")
	ErrInvalidEmail                = errors.New("email не валиден")
	ErrInvalidID                   = errors.New("id удаляемого пользователя не может быть равен 0")
)

// validateRequest - использовать для проверки получаемых запросов на ожидаемые данные
func validateRequest(req interface{}) error {
	v := reflect.ValueOf(req)

	// Проверка на указатель и получение значения
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Проверка, что входные данные являются структурой
	if v.Kind() != reflect.Struct {
		return errors.New("ожидалась структура для валидации")
	}

	// Итерация по полям структуры
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := v.Type().Field(i).Name

		switch fieldName {
		case email:
			if field.String() == "" {
				return ErrEmailNotSpecified
			}
			if !validator.IsValidEmail(field.String()) {
				return ErrInvalidEmail
			}
		case name:
			if field.String() == "" {
				return ErrNameNotSpecified
			}
		case password:
			if field.String() == "" {
				return ErrPasswordNotSpecified
			}
		case passwordConfirm:
			if field.String() == "" {
				return ErrPasswordConfirmNotSpecified
			}
		case role:
			roleID := field.Int()

			if roleID == int64(desc.Role_ROLE_UNSPECIFIED) {
				return ErrRoleNotSpecified
			} else if roleID != int64(desc.Role_ROLE_UNSPECIFIED) &&
				roleID != int64(desc.Role_ROLE_USER) &&
				roleID != int64(desc.Role_ROLE_ADMIN) {

				return ErrInvalidRole
			}
		case id:
			if field.Int() == 0 {
				return ErrInvalidID
			}
		}
	}
	return nil
}
