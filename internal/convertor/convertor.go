package convertor

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/erikqwerty/auth/internal/model"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

// ToCreateUserFromCreateRequest - Конвертер для преобразования gRPC-запроса создания пользователя в модель бизнес-логики User
func ToCreateUserFromCreateRequest(req *desc.CreateRequest) *model.CreateUser {
	if req == nil {
		return nil
	}
	return &model.CreateUser{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: req.Password,
		RoleID:       int32(req.Role),
	}
}

// ToUpdateUserFromUpdateRequest - Конвертор для преобразования gRPC-запроса обновления пользователя в модель бизнес-логики User
func ToUpdateUserFromUpdateRequest(req *desc.UpdateRequest) *model.UpdateUser {
	if req == nil {
		return nil
	}

	role := int32(req.Role)

	return &model.UpdateUser{
		Email:  &req.Email,
		Name:   &req.Name.Value,
		RoleID: &role,
	}
}

// ToGetResponseFromReadUser - Конвертор для преобразования модели бизнес-логики User в gRPC-ответ
func ToGetResponseFromReadUser(user *model.UserInfo) *desc.GetResponse {
	if user == nil {
		return nil
	}
	return &desc.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      desc.Role(user.RoleID),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
