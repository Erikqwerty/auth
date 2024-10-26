package api

import (
	"context"
	"errors"

	"github.com/erikqwerty/auth/internal/convertor"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
	"github.com/erikqwerty/auth/pkg/utils/validator"
)

// CreateUser - обрабатывает получаемый запрос от клиента gRPC на создание пользователя
func (i *ImplServAuthUser) CreateUser(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if err := validateDataCreateRequest(req); err != nil {
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

	if !validator.IsValidEmail(req.Email) {
		return errors.New("email не валиден")
	}

	return nil
}
