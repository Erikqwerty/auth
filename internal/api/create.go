package api

import (
	"context"
	"errors"
	"regexp"

	"github.com/erikqwerty/auth/internal/convertor"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// Create - обрабатывает получаемый запрос от клиента gRPC на создание пользователя
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("пароли не совпадают")
	}
	if !isValidEmail(req.Email) {
		return nil, errors.New("email не валиден")
	}

	id, err := i.authService.Create(ctx, convertor.ToModelUserFromCreateRequest(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: id}, err
}

// isValidEmail проверяет валидность email-адреса. Возвращает true если валидно.
func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
