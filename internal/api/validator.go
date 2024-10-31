package api

import (
	"errors"
	"reflect"

	"github.com/erikqwerty/auth/internal/autherrors"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
	"github.com/erikqwerty/auth/pkg/utils"
)

// валидируемые данные поля структур запроса gRPC контракта
const (
	email           = "Email"
	name            = "Name"
	password        = "Password"
	passwordConfirm = "PasswordConfirm"
	role            = "Role"
	id              = "ID"
)

// ValidateRequest - использовать для проверки получаемых запросов на ожидаемые данные
func ValidateRequest(req interface{}) error {
	v := reflect.ValueOf(req)

	// Проверка на указатель и получение значения
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Проверка, что входные данные являются структурой
	if v.Kind() != reflect.Struct {
		return errors.New("ожидалась структура для валидации")
	}

	// Рекурсивная проверка структуры на ожидаемые поля
	return validateStructFields(v)
}

// validateStructFields - рекурсивная функция для проверки полей структуры на наличие ожидаемых данных
func validateStructFields(v reflect.Value) error {
	var passwordValue, passwordConfirmValue string

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := v.Type().Field(i).Name

		// Проверка на вложенные структуры и рекурсивный вызов
		if field.Kind() == reflect.Struct {
			if err := validateStructFields(field); err != nil {
				return err
			}
			continue
		}

		// Проверка полей структуры на ожидаемые значения
		switch fieldName {
		case email:
			if field.String() == "" {
				return autherrors.ErrEmailNotSpecified
			}
			if !utils.IsValidEmail(field.String()) {
				return autherrors.ErrInvalidEmail
			}
		case name:
			if field.String() == "" {
				return autherrors.ErrNameNotSpecified
			}
		case password:
			if field.String() == "" {
				return autherrors.ErrPasswordNotSpecified
			}
			passwordValue = field.String()
		case passwordConfirm:
			if field.String() == "" {
				return autherrors.ErrPasswordConfirmNotSpecified
			}
			passwordConfirmValue = field.String()
		case role:
			roleID := field.Int()
			if roleID == int64(desc.Role_ROLE_UNSPECIFIED) {
				return autherrors.ErrRoleNotSpecified
			} else if roleID != int64(desc.Role_ROLE_USER) && roleID != int64(desc.Role_ROLE_ADMIN) {
				return autherrors.ErrInvalidRole
			}
		case id:
			if field.Int() == 0 {
				return autherrors.ErrInvalidID
			}
		}
	}

	if passwordValue != "" && passwordConfirmValue != "" && passwordValue != passwordConfirmValue {
		return autherrors.ErrPasswordsDoNotMatch
	}

	return nil
}
