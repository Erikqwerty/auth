package autherrors

import "errors"

// Ошибки для валидации
var (
	// не указан email
	ErrEmailNotSpecified = errors.New("не указан email")
	// имя пользователя указано пустым
	ErrNameNotSpecified = errors.New("имя пользователя указано пустым")
	// не указан пароль
	ErrPasswordNotSpecified = errors.New("не указан пароль")
	// не указан пароль подтверждения
	ErrPasswordConfirmNotSpecified = errors.New("не указан пароль подтверждения")
	// пароли не совпадают
	ErrPasswordsDoNotMatch = errors.New("пароли не совпадают")
	// роль пользователя не была указана
	ErrRoleNotSpecified = errors.New("роль пользователя не задан")
	// переданная роль пользователя не существует
	ErrInvalidRole = errors.New("переданная роль пользователя не существует")
	// email не валиден
	ErrInvalidEmail = errors.New("email не валиден")
	// id удаляемого пользователя не может быть равен 0
	ErrInvalidID = errors.New("id удаляемого пользователя не может быть равен 0")
)

var (
	ErrCreateUserNil = errors.New("invalid data")
)
