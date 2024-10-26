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
	switch {
	case req.Email == "":
		return errors.New("не указан email")

	case req.Name == "":
		return errors.New("нельзя оставлять имя пользователя пустым")

	case req.Password == "":
		return errors.New("не указан пароль")

	case req.PasswordConfirm == "":
		return errors.New("не указан пароль подтверждения")

	case req.Password != req.PasswordConfirm:
		return errors.New("пароли не совпадают")

	case req.Role == desc.Role_ROLE_UNSPECIFIED:
		return errors.New("роль пользователя не была указана")

	case req.Role != desc.Role_ROLE_ADMIN && req.Role != desc.Role_ROLE_USER:
		return errors.New("переданная роль пользователя не существует")

	case !validator.IsValidEmail(req.Email):
		return errors.New("email не валиден")

	default:
		return nil
	}
}
