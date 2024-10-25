package api

import (
	"context"
	"errors"
	"regexp"

	"github.com/erikqwerty/auth/internal/convertor"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// CreateUser - обрабатывает получаемый запрос от клиента gRPC на создание пользователя
func (i *ImplServAuthUser) CreateUser(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	err := validateDataCreateRequest(req)
	if err != nil {
		return nil, err
	}

	id, err := i.authService.CreateUser(ctx, convertor.ToCreateUserFromCreateRequest(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: id}, err
}

// validateDataCreateRequest - необходима для проверки переданных данных и их валидации перед обработкой в сервисном слое
func validateDataCreateRequest(req *desc.CreateRequest) error {
	if req.Email == "" {
		return errors.New("не указан email")
	}

	if req.Name == "" {
		return errors.New("не указано имя пользователя")
	}

	if req.Password == "" {
		return errors.New("не указан пароль")
	}

	if req.PasswordConfirm == "" {
		return errors.New("не указан пароль подтверждения")
	}

	if req.Password != req.PasswordConfirm {
		return errors.New("пароли не совпадают")
	}

	if !isValidEmail(req.Email) {
		return errors.New("email не валиден")
	}

	return nil
}

// isValidEmail проверяет валидность email-адреса. Возвращает true если валидно.
func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)

	return re.MatchString(email)
}
